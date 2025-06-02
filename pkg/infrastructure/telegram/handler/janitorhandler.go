package handler

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-telegram/bot/models"
	"github.com/nightnoryu/janitor/pkg/infrastructure/jsonlog"

	tg "github.com/go-telegram/bot"
)

func NewJanitorHandler(logger jsonlog.Logger) tg.HandlerFunc {
	return func(ctx context.Context, bot *tg.Bot, update *models.Update) {
		if update.Message.Chat.Type != models.ChatTypeSupergroup {
			return
		}
		if !strings.HasPrefix(update.Message.Text, "/ban") {
			return
		}
		if update.Message.ReplyToMessage == nil {
			return
		}

		_, err := bot.SendMessage(ctx, &tg.SendMessageParams{
			Text:   "ban user " + update.Message.ReplyToMessage.From.Username,
			ChatID: update.Message.Chat.ID,
		})
		if err != nil {
			logger.Error(err)
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
