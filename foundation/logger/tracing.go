package logger

import (
	"context"
)

// type 'ctxKey' represents a key for a value stored in a context
type ctxKey string

// constant 'defaultTraceIdKey' is the default key used to store the traceId in case of missing key
const defaultTraceIdKey ctxKey = "trace_id"

// constant 'defaultTraceIdValue' is the default value for the traceId in case of missing value
const defaultTraceIdValue = "default_trace_id"

// function 'getTraceId' returns the traceId from the context using the provided key,
// if the key is empty, it uses the default key defined as 'defaultTraceIdKey'
// if the value is missing, it returns the default value defined as 'defaultTraceIdValue'
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
