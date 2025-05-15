package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/nightnoryu/janitor/pkg/janitor/infrastructure/jsonlog"
	"github.com/nightnoryu/janitor/pkg/janitor/infrastructure/telegram/handler"

	"github.com/go-telegram/bot"
)

const appID = "janitor"

func main() {
	logger := initLogger()

	conf, err := parseEnv()
	if err != nil {
		logger.FatalError(err)
	}

	options := initBotOptions(logger)

	b, err := bot.New(conf.TelegramBotToken, options...)
	if err != nil {
		logger.FatalError(err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	b.Start(ctx)
}

func initLogger() jsonlog.Logger {
	logger := jsonlog.NewLogger(&jsonlog.Config{
		AppName: appID,
		Level:   jsonlog.InfoLevel,
	})
	return logger
}

func initBotOptions(logger jsonlog.Logger) []bot.Option {
	janitorHandler := handler.NewJanitorHandler(logger)
	return []bot.Option{
		bot.WithDefaultHandler(janitorHandler),
	}
}
