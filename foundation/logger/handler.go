package logger

import (
	"context"
)

// type 'Handler' represents a logging handler,
// concrete implementations are 'ConsoleHandler' and 'FileHandler'
type Handler interface {
	Handle(ctx context.Context, r Record)
}
