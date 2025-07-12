package logger

import "context"

// type 'Hook' represents a logging hook
type Hook interface {
	// function 'OnAll' is called on each and every record
	OnAll(ctx context.Context, r Record)
	// function 'OnDebug' is called on debug level records
	OnDebug(ctx context.Context, r Record)
	// function 'OnInfo' is called on info level records
	OnInfo(ctx context.Context, r Record)
	// function 'OnWarn' is called on warning level records
	OnWarn(ctx context.Context, r Record)
	// function 'OnError' is called on error level records
	OnError(ctx context.Context, r Record)
}
