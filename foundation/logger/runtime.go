package logger

import (
	"runtime"
)

type RuntimeInfo struct {
	GoVersion    string
	GoMaxProcs   int
	NumCPU       int
	NumGoroutine int
}

func NewRuntimeInfo() RuntimeInfo {
	return RuntimeInfo{
		GoVersion:    runtime.Version(),
		GoMaxProcs:   runtime.GOMAXPROCS(0),
		NumCPU:       runtime.NumCPU(),
		NumGoroutine: runtime.NumGoroutine(),
	}
}

func (r RuntimeInfo) Fields() []Field {
	return []Field{
		{"go_version", r.GoVersion},
		{"go_maxprocs", r.GoMaxProcs},
		{"num_cpu", r.NumCPU},
		{"num_goroutine", r.NumGoroutine},
	}
}
