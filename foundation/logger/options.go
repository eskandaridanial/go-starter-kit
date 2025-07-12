package logger

import "context"

type Option func(*Logger)

// function 'WithLevel' returns an option to set minimum logging level
func WithLevel(level Level) Option {
	return func(l *Logger) {
		l.Level = level
	}
}

// function 'WithHandler' returns an option to add a handler to the logger
func WithHandler(h Handler) Option {
	return func(l *Logger) {
		l.Handlers = append(l.Handlers, h)
	}
}

// function 'WithField' returns an option to add a field to the logger
func WithField(f Field) Option {
	return func(l *Logger) {
		l.Fields = append(l.Fields, f)
	}
}

// function 'WithHook' returns an option to add a hook to the logger
func WithHook(h Hook) Option {
	return func(l *Logger) {
		l.Hooks = append(l.Hooks, h)
	}
}

// function 'WithContext' returns an option to set the context for the logger
func WithContext(ctx context.Context) Option {
	return func(l *Logger) {
		l.ctx = ctx
	}
}

// function 'WithBuildInfo' returns an option to add build information to the logger
func WithBuildInfo(versionKey, commitKey, timeKey string, once bool) Option {
	return func(l *Logger) {
		info := NewBuildInfo(versionKey, commitKey, timeKey)
		if once {
			l.onceBuildInfo = info.Fields()
		} else {
			newFields := make([]Field, 0, len(l.Fields))
			newFields = append(newFields, l.Fields...)
			l.Fields = append(newFields, info.Fields()...)
		}
	}
}

// function 'WithRuntimeInfo' returns an option to add runtime information to the logger
func WithRuntimeInfo(once bool) Option {
	return func(l *Logger) {
		info := NewRuntimeInfo()
		if once {
			l.onceRuntimeInfo = info.Fields()
		} else {
			newFields := make([]Field, 0, len(l.Fields))
			newFields = append(newFields, l.Fields...)
			l.Fields = append(newFields, info.Fields()...)
		}
	}
}

// function 'WithService' returns an option to add service name to the logger as a field
func WithService(service string) Option {
	return func(l *Logger) {
		l.Fields = append(l.Fields, String("service", service))
	}
}

// function 'WithServiceEnv' returns an option to add service name to the logger as a field,
// the service name is retrieved from the environment variable with the given key,
// if the environment variable is not set, it returns the key name itself
func WithServiceEnv(service string) Option {
	return func(l *Logger) {
		l.Fields = append(l.Fields, String("service", getEnvOrKey(service)))
	}
}

// function 'WithEnvironment' returns an option to add environment name to the logger as a field
func WithEnvironment(environment string) Option {
	return func(l *Logger) {
		l.Fields = append(l.Fields, String("environment", environment))
	}
}

// function 'WithEnvironmentEnv' returns an option to add environment name to the logger as a field,
// the environment name is retrieved from the environment variable with the given key,
// if the environment variable is not set, it returns the key name itself
func WithEnvironmentEnv(environment string) Option {
	return func(l *Logger) {
		l.Fields = append(l.Fields, String("environment", getEnvOrKey(environment)))
	}
}

// function 'WithTraceIdKey' returns an option to set the traceId key for the logger,
// the traceId key is used to retrieve the traceId value from the context
func WithTraceIdKey(traceIdKey string) Option {
	return func(l *Logger) {
		l.traceIdKey = traceIdKey
	}
}

// function 'WithBufferSize' returns an option to set the buffer size for the logger,
// the buffer size is used to limit the number of log records that can be buffered
func WithBufferSize(size int) Option {
	return func(l *Logger) {
		l.bufferSize = size
	}
}

// function 'WithBackpressure' returns an option to set the backpressure strategy for the logger,
// the backpressure strategy is used to control the behavior of the logger when the buffer is full
func WithBackpressure(strategy BackpressureStrategy) Option {
	return func(l *Logger) {
		l.backpressure = strategy
	}
}

// function 'WithWorkers' returns an option to set the number of workers for the logger,
// the number of workers is used to control the number of worker goroutines that process log records
func WithWorkers(n int) Option {
	return func(l *Logger) {
		l.numWorkers = n
	}
}

// function 'WithInternalErrorHandler' returns an option to set the internal error handler for the logger,
// the internal error handler is used to handle errors that occur within the logger
func WithInternalErrorHandler(f func(error)) Option {
	return func(l *Logger) {
		l.internalErrorHandler = f
	}
}
