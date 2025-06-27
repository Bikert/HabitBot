package utils

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func GetMessage(update *tgbotapi.Update) *tgbotapi.Message {
	if update.CallbackQuery != nil {
		return update.CallbackQuery.Message
	}
	return update.Message
}

func GetUserId(update *tgbotapi.Update) *tgbotapi.User {
	if update.CallbackQuery != nil {
		return update.CallbackQuery.From
	}
	return update.Message.From
}
