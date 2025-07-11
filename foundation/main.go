package main

import (
	"context"

	"github.com/eskandaridanial/go-starter-kit/foundation/logger"
	"github.com/eskandaridanial/go-starter-kit/foundation/logger/handlers"
	"github.com/eskandaridanial/go-starter-kit/foundation/logger/hooks"
)

func main() {
	log := logger.NewLogger(
		logger.WithLevel(logger.Debug),
		logger.WithHandler(&handlers.ConsoleHandler{Formatter: logger.NewJSONFormatter()}),
		logger.WithField(logger.String("key", "value")),
		logger.WithHook(&hooks.DefaultHook{}),
		logger.WithContext(context.Background()),
		logger.WithBuildInfo("OS_ENV_FOR_VERSION", "OS_ENV_FOR_COMMIT", "OS_ENV_FOR_TIME", true),
		logger.WithRuntimeInfo(true),
		logger.WithService("auth"),
		logger.WithServiceEnv("OS_ENV_FOR_SERVICE_ENV"),
		logger.WithEnvironment("dev"),
		logger.WithEnvironmentEnv("OS_ENV_FOR_ENV"),
		logger.WithTraceIdKey("OS_ENV_FOR_TRACE_ID_KEY"),
	)
	defer log.Close()

	log.Debug("user logged in")
	log.Debug("user logged in", logger.String("key", "value"))
	log.DebugCtx(context.Background(), "user logged in")
	log.DebugCtx(context.Background(), "user logged in", logger.String("key", "value"))

	log.Info("user logged in")
	log.Info("user logged in", logger.String("key", "value"))
	log.InfoCtx(context.Background(), "user logged in")
	log.InfoCtx(context.Background(), "user logged in", logger.String("key", "value"))

	log.Error("user logged in")
	log.Error("user logged in", logger.String("key", "value"))
	log.ErrorCtx(context.Background(), "user logged in")
	log.ErrorCtx(context.Background(), "user logged in", logger.String("key", "value"))

	log.Warn("user logged in")
	log.Warn("user logged in", logger.String("key", "value"))
	log.WarnCtx(context.Background(), "user logged in")
	log.WarnCtx(context.Background(), "user logged in", logger.String("key", "value"))
}
