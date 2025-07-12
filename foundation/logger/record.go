package logger

import "time"

// type 'Record' represents a log record for structured logging
type Record struct {
	Level     Level
	Message   string
	TraceId   string
	Caller    string
	Fields    []Field
	Timestamp time.Time
}
