package handlers

import (
	"context"
	"os"

	"github.com/eskandaridanial/go-starter-kit/foundation/logger"
)

// struct 'ConsoleHandler' implements 'Handler' interface
type ConsoleHandler struct {
	Formatter logger.Formatter
}

// function 'Handle' handles the given record by formatting it and writing it to the console
func (h *ConsoleHandler) Handle(ctx context.Context, r logger.Record) {
	output := h.Formatter.Format(r)
	os.Stdout.Write(output)
}
