package hooks

import (
	"context"

	"github.com/eskandaridanial/go-starter-kit/foundation/logger"
)

// struct 'DefaultHook' implements 'Hook' interface
type DefaultHook struct{}

// function 'OnAll' is called on each and every record
func (h *DefaultHook) OnAll(ctx context.Context, r logger.Record) {}

// function 'OnDebug' is called on debug level records
func (h *DefaultHook) OnDebug(ctx context.Context, r logger.Record) {}

// function 'OnInfo' is called on info level records
func (h *DefaultHook) OnInfo(ctx context.Context, r logger.Record) {}

// function 'OnWarn' is called on warning level records
func (h *DefaultHook) OnWarn(ctx context.Context, r logger.Record) {}

// function 'OnError' is called on error level records
func (h *DefaultHook) OnError(ctx context.Context, r logger.Record) {}
