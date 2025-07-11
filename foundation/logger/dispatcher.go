package logger

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var defaultBufferSize = 1000

type Dispatcher struct {
	records  chan dispatchEntry
	wg       sync.WaitGroup
	handlers []Handler
	hooks    []Hook
}

type dispatchEntry struct {
	ctx context.Context
	rec Record
}

func NewDispatcher(handlers []Handler, hooks []Hook) *Dispatcher {
	d := &Dispatcher{
		records:  make(chan dispatchEntry, defaultBufferSize),
		handlers: handlers,
		hooks:    hooks,
	}
	d.wg.Add(1)
	go d.run()
	return d
}

func (d *Dispatcher) Dispatch(ctx context.Context, rec Record) {
	select {
	case d.records <- dispatchEntry{ctx: ctx, rec: rec}:
	default:
		fmt.Printf("logger: dropped log due to full queue: %s", rec.Message)
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
			hook.OnAll(ctx, rec)
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
		fmt.Printf("logger: recovered from panic: %v", r)
	}
}
