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
		if len(update.Message.NewChatMembers) > 0 && update.Message.NewChatMembers[0].ID == bot.ID() {
			err := handleBotAddedToChat(ctx, bot, update)
			if err != nil {
				logger.Error(err)
			}
			return
		}

		if update.Message.Chat.Type == models.ChatTypePrivate {
			err := handlePrivateChatMessage(ctx, bot, update)
			if err != nil {
				logger.Error(err)
			}
			return
		}

		if update.Message.Chat.Type != models.ChatTypeSupergroup {
			return
		}

		if strings.HasPrefix(update.Message.Text, "/ban") {
			err := handleBanMember(ctx, bot, update)
			if err != nil {
				logger.Error(err)
			}
		}
	}
}

func handleBotAddedToChat(ctx context.Context, bot *tg.Bot, update *models.Update) error {
	text := botAddedToChatMessage

	chatMember, err := bot.GetChatMember(ctx, &tg.GetChatMemberParams{
		ChatID: update.Message.Chat.ID,
		UserID: bot.ID(),
	})
	if err != nil {
		return err
	}

	if chatMember.Administrator == nil {
		text += adminNotAssignedMessage
	} else {
		text += adminAssignedMessage
	}

	_, err = bot.SendMessage(ctx, &tg.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   text,
	})
	return err
}

func handlePrivateChatMessage(ctx context.Context, bot *tg.Bot, update *models.Update) error {
	_, err := bot.SendMessage(ctx, &tg.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   privateChatMessage,
	})
	return err
}

func handleBanMember(ctx context.Context, bot *tg.Bot, update *models.Update) error {
	reply := update.Message.ReplyToMessage
	if reply == nil {
		return nil
	}

	_, err := bot.BanChatMember(ctx, &tg.BanChatMemberParams{
		ChatID: update.Message.Chat.ID,
		UserID: reply.From.ID,
	})
	if err != nil {
		return err
	}

	_, err = bot.SendMessage(ctx, &tg.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   fmt.Sprintf(memberBannedMessageTemplate, reply.From.Username),
		ReplyParameters: &models.ReplyParameters{
			MessageID: update.Message.ReplyToMessage.ID,
		},
	})
	if err != nil {
		return err
	}

	_, err = bot.DeleteMessage(ctx, &tg.DeleteMessageParams{
		ChatID:    update.Message.Chat.ID,
		MessageID: update.Message.ReplyToMessage.ID,
	})

	return err
}
