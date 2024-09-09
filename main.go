package main

import (
	"dexpert-event-listener/listener"
	"log/slog"

	"github.com/x1rh/event-listener/logger"
)

func main() {
	logger.Init(slog.LevelInfo, false)
	el, err := listener.NewTokenFactoryEventListener()
	if err != nil {
		slog.Error("fail to new a token factory listener", slog.Any("err", err))
	}
	el.Start()
}
