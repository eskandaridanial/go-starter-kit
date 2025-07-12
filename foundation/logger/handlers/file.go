package handlers

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/eskandaridanial/go-starter-kit/foundation/logger"
)

// struct 'FileHandler' implements 'Handler' interface
type FileHandler struct {
	File      *os.File
	Formatter logger.Formatter
	mu        sync.Mutex
}

// function 'Handle' handles the given record by formatting it and writing it to the file
func (h *FileHandler) Handle(ctx context.Context, r logger.Record) {
	h.mu.Lock()
	defer h.mu.Unlock()
	output := h.Formatter.Format(r)
	_, err := h.File.Write(output)
	if err != nil {
		fmt.Printf("logger: file write error: %v\n", err)
	}
}
