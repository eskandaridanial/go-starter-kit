package handlers

import (
	"context"
	"os"

	"github.com/eskandaridanial/go-starter-kit/foundation/logger"
)

type ConsoleHandler struct {
	Formatter logger.Formatter
}

func (h *ConsoleHandler) Handle(ctx context.Context, r logger.Record) {
	output := h.Formatter.Format(r)
	os.Stdout.Write(output)
}
