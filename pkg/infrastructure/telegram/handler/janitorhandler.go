package handler

import (
	"context"
	"fmt"

	"github.com/go-telegram/bot/models"
	"github.com/nightnoryu/janitor/pkg/infrastructure/jsonlog"

	tg "github.com/go-telegram/bot"
)

func NewJanitorHandler(logger jsonlog.Logger) tg.HandlerFunc {
	return func(ctx context.Context, bot *tg.Bot, update *models.Update) {
		if update.Message == nil {
			return
		}

		logger.Info(fmt.Sprintf(
			"message from %s in chat %s (ID=%d): %s",
			update.Message.From.Username,
			update.Message.Chat.Title,
			update.Message.Chat.ID,
			update.Message.Text,
		))
	}
}
