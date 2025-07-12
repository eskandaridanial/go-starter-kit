package logger

import (
	"runtime"
)

// type 'RuntimeInfo' represents the runtime information of the system
type RuntimeInfo struct {
	GoVersion    string
	GoMaxProcs   int
	NumCPU       int
	NumGoroutine int
}

// function 'NewRuntimeInfo' returns a new instance of 'RuntimeInfo' with the current runtime information
func NewRuntimeInfo() RuntimeInfo {
	return RuntimeInfo{
		GoVersion:    runtime.Version(),
		GoMaxProcs:   runtime.GOMAXPROCS(0),
		NumCPU:       runtime.NumCPU(),
		NumGoroutine: runtime.NumGoroutine(),
	}
}

// function 'Fields' returns the runtime information as a slice of 'Field' for logging
func (r RuntimeInfo) Fields() []Field {
	return []Field{
		{"go_version", r.GoVersion},
		{"go_maxprocs", r.GoMaxProcs},
		{"num_cpu", r.NumCPU},
		{"num_goroutine", r.NumGoroutine},
	}
}
