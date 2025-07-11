package logger

import (
	"context"
)

type Handler interface {
	Handle(ctx context.Context, r Record)
}
