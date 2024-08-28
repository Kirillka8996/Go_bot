package main

import (
	"Go_bot/config"
	"Go_bot/pkg/handlers"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	token := config.BOT_TOKEN
	if token == "" {
		log.Fatal("Bot Token is empty")
	}
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("Error creating bot: %v", err)
	}
	bot.Debug = true
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		go handlers.HandleUpdate(bot, update) // Обработка каждого обновления в отдельной горутине
	}
}
