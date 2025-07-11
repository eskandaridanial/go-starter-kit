package logger

import "context"

type Option func(*Logger)

func WithLevel(level Level) Option {
	return func(l *Logger) {
		l.Level = level
	}
}

func WithHandler(h Handler) Option {
	return func(l *Logger) {
		l.Handlers = append(l.Handlers, h)
	}
}

func WithField(f Field) Option {
	return func(l *Logger) {
		l.Fields = append(l.Fields, f)
	}
}

func WithHook(h Hook) Option {
	return func(l *Logger) {
		l.Hooks = append(l.Hooks, h)
	}
}

func WithContext(ctx context.Context) Option {
	return func(l *Logger) {
		l.ctx = ctx
	}
}

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

func WithService(service string) Option {
	return func(l *Logger) {
		l.Fields = append(l.Fields, String("service", service))
	}
}

func WithServiceEnv(service string) Option {
	return func(l *Logger) {
		l.Fields = append(l.Fields, String("service", getEnvOrKey(service)))
	}
}

func WithEnvironment(environment string) Option {
	return func(l *Logger) {
		l.Fields = append(l.Fields, String("environment", environment))
	}
}

func WithEnvironmentEnv(environment string) Option {
	return func(l *Logger) {
		l.Fields = append(l.Fields, String("environment", getEnvOrKey(environment)))
	}
}

func WithTraceIdKey(traceIdKey string) Option {
	return func(l *Logger) {
		l.traceIdKey = traceIdKey
	}
}
