package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/nightnoryu/janitor/pkg/infrastructure/jsonlog"
	"github.com/nightnoryu/janitor/pkg/infrastructure/telegram/handler"
	"github.com/nightnoryu/janitor/pkg/infrastructure/telegram/initializer"

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

	err = initializer.InitializeCommands(ctx, b)
	if err != nil {
		logger.FatalError(err)
	}

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
