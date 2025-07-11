package logger

import (
	"context"
	"fmt"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

type BackpressureStrategy int

const (
	Drop BackpressureStrategy = iota
	Block
)

type Dispatcher struct {
	records             chan dispatchEntry
	wg                  sync.WaitGroup
	handlers            []Handler
	hooks               []Hook
	numWorkers          int
	backpressure        BackpressureStrategy
	bufferSize          int
	dropNoticeThreshold int64

	internalErrorHandler func(error)

	droppedLogsCount int64
}

type dispatchEntry struct {
	ctx context.Context
	rec Record
}

func NewDispatcher(
	handlers []Handler,
	hooks []Hook,
	numWorkers int,
	bufferSize int,
	backpressure BackpressureStrategy,
	internalErrorHandler func(error),
) *Dispatcher {
	if numWorkers <= 0 {
		numWorkers = 1
	}
	if bufferSize <= 0 {
		bufferSize = 1000
	}

	d := &Dispatcher{
		records:              make(chan dispatchEntry, bufferSize),
		handlers:             handlers,
		hooks:                hooks,
		numWorkers:           numWorkers,
		backpressure:         backpressure,
		bufferSize:           bufferSize,
		dropNoticeThreshold:  1000,
		internalErrorHandler: internalErrorHandler,
	}

	for i := 0; i < numWorkers; i++ {
		d.wg.Add(1)
		go d.run()
	}

	return d
}

func (d *Dispatcher) Dispatch(ctx context.Context, rec Record) {
	entry := dispatchEntry{ctx: ctx, rec: rec}

	switch d.backpressure {
	case Drop:
		select {
		case d.records <- entry:
		default:
			atomic.AddInt64(&d.droppedLogsCount, 1)
			if d.DroppedCount()%d.dropNoticeThreshold == 0 {
				d.reportInternalError(fmt.Errorf("dropped %d logs due to full queue", d.DroppedCount()))
			}
		}
	case Block:
		d.records <- entry
	}
}

func (d *Dispatcher) Close() {
	close(d.records)
	d.wg.Wait()
}

func (d *Dispatcher) run() {
	defer d.wg.Done()
	for entry := range d.records {
		d.deliver(entry.ctx, entry.rec)
	}
}

func (d *Dispatcher) deliver(ctx context.Context, rec Record) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	for _, h := range d.handlers {
		func() {
			defer d.recover()
			h.Handle(ctx, rec)
		}()
	}

	for _, hook := range d.hooks {
		func() {
			defer d.recover()
			switch rec.Level {
			case Debug:
				hook.OnDebug(ctx, rec)
			case Info:
				hook.OnInfo(ctx, rec)
			case Warn:
				hook.OnWarn(ctx, rec)
			case Error:
				hook.OnError(ctx, rec)
			}
		}()
	}
}

func (d *Dispatcher) recover() {
	if r := recover(); r != nil {
		d.reportInternalError(fmt.Errorf("recovered from panic in handler/hook: %v", r))
	}
}

func (d *Dispatcher) DroppedCount() int64 {
	return atomic.LoadInt64(&d.droppedLogsCount)
}

func (d *Dispatcher) BufferSize() int {
	return d.bufferSize
}

func (d *Dispatcher) reportInternalError(err error) {
	if d.internalErrorHandler != nil {
		d.internalErrorHandler(err)
	} else {
		internalLog("%v", err)
	}
}

func internalLog(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "logger [internal]: "+format+"\n", args...)
}
