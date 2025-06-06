package initializer

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func InitializeCommands(ctx context.Context, b *bot.Bot) error {
	_, err := b.SetMyCommands(ctx, &bot.SetMyCommandsParams{
		Commands: []models.BotCommand{
			{
				Command:     "ban",
				Description: "[ответ] забанить",
			},
		},
		Scope: &models.BotCommandScopeAllChatAdministrators{},
	})
	if err != nil {
		return err
	}

	_, err = b.SetMyCommands(ctx, &bot.SetMyCommandsParams{
		Commands: []models.BotCommand{},
		Scope:    &models.BotCommandScopeAllPrivateChats{},
	})
	if err != nil {
		return err
	}

	return nil
}
