package hooks

import (
	"context"

	"github.com/eskandaridanial/go-starter-kit/foundation/logger"
)

type DefaultHook struct{}

func (h *DefaultHook) OnAll(ctx context.Context, r logger.Record) {}

func (h *DefaultHook) OnDebug(ctx context.Context, r logger.Record) {}

func (h *DefaultHook) OnInfo(ctx context.Context, r logger.Record) {}

func (h *DefaultHook) OnWarn(ctx context.Context, r logger.Record) {}

func (h *DefaultHook) OnError(ctx context.Context, r logger.Record) {}
