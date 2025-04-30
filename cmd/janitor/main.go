package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/nightnoryu/janitor/pkg/infrastructure/jsonlog"

	"github.com/go-telegram/bot"
)

const appID = "janitor"

func main() {
	logger := initLogger()

	conf, err := parseEnv()
	if err != nil {
		logger.FatalError(err)
	}

	opts := []bot.Option{
		bot.WithDebug(),
	}

	b, err := bot.New(conf.TelegramBotToken, opts...)
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
