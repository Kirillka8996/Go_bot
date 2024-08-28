package handlers

import (
	"Go_bot/config"
	"Go_bot/pkg/basicfunc"
	"Go_bot/pkg/httpjson"
	"Go_bot/pkg/keyboard"
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// импортируем клавиатуру для игры КНБ
var gameKeyboard = keyboard.GameKeyboard

func HandleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	// Создаем сообщение для ответа
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	// Обрабатываем команды
	if update.Message.IsCommand() {
		handleCommand(&msg, update)
	} else {
		handleTextMessage(bot, &msg, update)
	}

	// Если есть текст для отправки, отправляем его
	if msg.Text != "" {
		basicfunc.SendMessage(bot, msg)
	}

	// Отправляем сообщение в группу
	basicfunc.SendGroupMessage(bot, update)
}

func handleCommand(msg *tgbotapi.MessageConfig, update tgbotapi.Update) {
	switch update.Message.Command() {
	case "start":
		msg.Text = fmt.Sprintf("Привет %s %s\nЯ бот работающий на Esp32", update.Message.From.FirstName, update.Message.From.LastName)
	case "weather":
		handleWeather(msg)
	case "random_insult":
		handleRandomInsult(msg)
	case "how_you_gosling":
		msg.Text = fmt.Sprintf("Ты Goслинг на %s%%", basicfunc.HowYouGosling())
	case "exchange_rates":
		handleExchangeRates(msg)
	case "game":
		msg.ReplyMarkup = gameKeyboard
		msg.Text = "Камень, Ножницы, Бумага"
	}
}

func handleTextMessage(bot *tgbotapi.BotAPI, msg *tgbotapi.MessageConfig, update tgbotapi.Update) {
	switch strings.ToLower(update.Message.Text) {
	case "привет":
		msg.Text = fmt.Sprintf("Привет %s %snЯ бот", update.Message.From.FirstName, update.Message.From.LastName)
	case "закрыть меню":
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		msg.Text = "Меню закрыто"
	case "🪨":
		msg.Text = basicfunc.Game(bot, update, 0)
	case "✂️":
		msg.Text = basicfunc.Game(bot, update, 1)
	case "📄":
		msg.Text = basicfunc.Game(bot, update, 2)
	}
}

func handleWeather(msg *tgbotapi.MessageConfig) {
	weatherToken := config.WEATER_TOKEN
	if weatherToken == "" {
		log.Fatal("Weather Token is empty")
	}
	data, err := httpjson.FetchJsonFromUrl(fmt.Sprintf("http://api.openweathermap.org/data/2.5/find?q=Ufa&type=like&appid=%s", weatherToken))
	if err != nil {
		msg.Text = "Ошибка при получении данных о погоде: "
	} else {
		msg.Text = basicfunc.Weather(data)
	}
}

func handleRandomInsult(msg *tgbotapi.MessageConfig) {
	data, err := httpjson.FetchJsonFromUrl("https://evilinsult.com/generate_insult.php?lang=ru&type=json")
	if err != nil {
		msg.Text = "Ошибка при получении оскорбления: " + err.Error()
	} else {
		msg.Text = basicfunc.RandomInsult(data)
	}
}

func handleExchangeRates(msg *tgbotapi.MessageConfig) {
	data, err := httpjson.FetchJsonFromUrl("https://www.cbr-xml-daily.ru/daily_json.js")
	if err != nil {
		msg.Text = "Ошибка при получении курса валют: " + err.Error()
	} else {
		msg.Text = basicfunc.ExchangeRates(data)
	}
}

// func recoverFromlog.Println() {
// 	//функция которая восстанавливает управление при критических ошибках и перезапускает программу
// 	if r := recover(); r != nil {
// 		fmt.Println("Recovered from log.Println:", r)
// 		fmt.Println("Перезапуск программы...")
// 		cmd := exec.Command(os.Args[0])
// 		cmd.Stdout = os.Stdout
// 		cmd.Stderr = os.Stderr
// 		cmd.Start()
// 		os.Exit(1)
// 	}
// }
