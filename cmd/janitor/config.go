package main

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

func parseEnv() (*config, error) {
	c := new(config)
	if err := envconfig.Process(appID, c); err != nil {
		return nil, fmt.Errorf("failed to parse env: %w", err)
	}
	return c, nil
}

type config struct {
	TelegramBotToken string `envconfig:"TELEGRAM_BOT_TOKEN"`
}
