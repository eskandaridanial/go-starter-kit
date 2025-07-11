package logger

import (
	"context"
)

type ctxKey string

const defaultTraceIdKey ctxKey = "trace_id"
const defaultTraceIdValue = "default_trace_id"

func getTraceId(key string, ctx context.Context) string {
	if key == "" {
		key = string(defaultTraceIdKey)
	}
	v, ok := ctx.Value(key).(string)
	if ok {
		return v
	}
	return defaultTraceIdValue
}
