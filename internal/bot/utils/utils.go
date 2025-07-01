package utils

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

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
func СonfirmPressAndHideButtons(api *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) error {
	callbackConfig := tgbotapi.NewCallback(callback.ID, "")
	_, err := api.Request(callbackConfig)
	if err != nil {
		return err
	}

	edit := tgbotapi.NewEditMessageTextAndMarkup(
		callback.Message.Chat.ID,
		callback.Message.MessageID,
		"⏳ Loading...",
		tgbotapi.InlineKeyboardMarkup{InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{}},
	)

	if _, err := api.Request(edit); err != nil {
		log.Println("failed to edit message:", err)
		return err
	}
	return nil
}
