package logger

import (
	"context"
	"time"
)

// type 'Logger' represents a message that contains properties required for structured logging
type Logger struct {
	Level                Level
	Handlers             []Handler
	Fields               []Field
	Hooks                []Hook
	ctx                  context.Context
	dispatcher           *Dispatcher
	onceBuildInfo        []Field
	onceRuntimeInfo      []Field
	traceIdKey           string
	bufferSize           int
	backpressure         BackpressureStrategy
	numWorkers           int
	internalErrorHandler func(error)
}

// function 'NewLogger' creates a new logger instance with the given options
func NewLogger(opts ...Option) *Logger {
	l := &Logger{
		Level: Info,
	}
	for _, o := range opts {
		o(l)
	}

	l.dispatcher = NewDispatcher(l.Handlers, l.Hooks, l.numWorkers, l.bufferSize, l.backpressure, l.internalErrorHandler)

	if len(l.onceBuildInfo) > 0 {
		l.Info("build information", l.onceBuildInfo...)
	}
	l.onceBuildInfo = nil

	if len(l.onceRuntimeInfo) > 0 {
		l.Info("runtime information", l.onceRuntimeInfo...)
	}
	l.onceRuntimeInfo = nil

	return l
}

// function 'WithField' creates a new logger instance with the given field added to the logger
// it is used after 'Logger' is initialized to add additional fields to the logger
func (l *Logger) WithField(field Field) *Logger {
	newFields := make([]Field, len(l.Fields)+1)
	copy(newFields, l.Fields)
	newFields[len(l.Fields)] = field

	return &Logger{
		Level:      l.Level,
		Handlers:   l.Handlers,
		Fields:     newFields,
		Hooks:      l.Hooks,
		ctx:        l.ctx,
		dispatcher: l.dispatcher,
	}
}

// function 'log' logs a message with the given level and fields
// it is used by 'Debug', 'Info', 'Warn', 'Error' methods
func (l *Logger) log(ctx context.Context, level Level, msg string, fields []Field) {
	if level < l.Level {
		return
	}

	var workingCtx context.Context
	if ctx != nil && ctx != context.TODO() && ctx != context.Background() {
		workingCtx = ctx
	} else if l.ctx != nil {
		workingCtx = l.ctx
	} else {
		workingCtx = ctx
	}

	merged := make([]Field, 0, len(l.Fields)+len(fields))
	merged = append(merged, l.Fields...)
	merged = append(merged, fields...)

	rec := Record{
		Level:     level,
		Message:   msg,
		TraceId:   getTraceId(l.traceIdKey, workingCtx),
		Caller:    caller(3),
		Fields:    merged,
		Timestamp: time.Now(),
	}

	l.dispatcher.Dispatch(workingCtx, rec)
}

// function 'Debug' logs a debug message with the given fields
func (l *Logger) Debug(msg string, fields ...Field) {
	l.log(context.Background(), Debug, msg, fields)
}

// function 'Info' logs an info message with the given fields
func (l *Logger) Info(msg string, fields ...Field) {
	l.log(context.Background(), Info, msg, fields)
}

// function 'Warn' logs a warning message with the given fields
func (l *Logger) Warn(msg string, fields ...Field) {
	l.log(context.Background(), Warn, msg, fields)
}

// function 'Error' logs an error message with the given fields
func (l *Logger) Error(msg string, fields ...Field) {
	l.log(context.Background(), Error, msg, fields)
}

// function 'DebugCtx' logs a debug message with the given fields and context
func (l *Logger) DebugCtx(ctx context.Context, msg string, fields ...Field) {
	l.log(ctx, Debug, msg, fields)
}

// function 'InfoCtx' logs an info message with the given fields and context
func (l *Logger) InfoCtx(ctx context.Context, msg string, fields ...Field) {
	l.log(ctx, Info, msg, fields)
}

// function 'WarnCtx' logs a warning message with the given fields and context
func (l *Logger) WarnCtx(ctx context.Context, msg string, fields ...Field) {
	l.log(ctx, Warn, msg, fields)
}

// function 'ErrorCtx' logs an error message with the given fields and context
func (l *Logger) ErrorCtx(ctx context.Context, msg string, fields ...Field) {
	l.log(ctx, Error, msg, fields)
}

// function 'Close' closes the logger and releases all resources
func (l *Logger) Close() {
	l.dispatcher.Close()
}
