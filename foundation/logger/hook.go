package logger

import "context"

type Hook interface {
	OnAll(ctx context.Context, r Record)
	OnDebug(ctx context.Context, r Record)
	OnInfo(ctx context.Context, r Record)
	OnWarn(ctx context.Context, r Record)
	OnError(ctx context.Context, r Record)
}
