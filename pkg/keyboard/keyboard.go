package keyboard

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var GameKeyboard = tgbotapi.NewReplyKeyboard(

	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🪨"),
		tgbotapi.NewKeyboardButton("✂️"),
		tgbotapi.NewKeyboardButton("📄"),
	),

	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Закрыть меню"),
	),
)
