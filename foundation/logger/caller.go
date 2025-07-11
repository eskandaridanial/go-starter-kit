package logger

import (
	"fmt"
	"runtime"
)

func caller(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "unknown:0"
	}

	short := file
	for i := len(file) - 1; i >= 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			break
		}
	}

	return fmt.Sprintf("%s:%d", short, line)
}
