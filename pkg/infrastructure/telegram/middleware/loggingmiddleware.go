package middleware

import (
	"context"

	"github.com/nightnoryu/janitor/pkg/infrastructure/jsonlog"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

const (
	chatIDField   = "chat_id"
	chatTypeField = "chat_type"
	usernameField = "username"
)

func NewLoggingMiddleware(logger jsonlog.Logger) bot.Middleware {
	return func(next bot.HandlerFunc) bot.HandlerFunc {
		return func(ctx context.Context, bot *bot.Bot, update *models.Update) {
			if update.Message == nil {
				return
			}

			chatLogger := logger.
				WithField(chatIDField, update.Message.Chat.ID).
				WithField(chatTypeField, update.Message.Chat.Type).
				WithField(usernameField, update.Message.From.Username)

			text := update.Message.Text
			if len(update.Message.Caption) > 0 {
				text = update.Message.Caption
			}

			chatLogger.Info(text)

			next(ctx, bot, update)
		}
	}
}
