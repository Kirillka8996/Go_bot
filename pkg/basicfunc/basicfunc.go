package basicfunc

import (
	"Go_bot/pkg/convert"
	"fmt"
	"log"
	"math/rand"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// функционал команд бота
func Weather(data map[string]any) (message string) {
	message = "Погода от Райана GOслинга\nТекущая погода в Уфе:\n\n"

	list, ok := data["list"].([]any)
	if !ok {
		message += "Ошибка: невозможно получить данные о погоде\n"
		return message
	}

	if len(list) > 0 {
		firstItem, ok := list[0].(map[string]any)
		if !ok {
			message += "Ошибка: невозможно получить данные о погоде\n"
			return message
		}

		mainData, ok := firstItem["main"].(map[string]any)
		if !ok {
			message += "Ошибка: невозможно получить данные о погоде\n"
			return message
		}

		tempKelvin, ok := mainData["temp"].(float64)
		if !ok {
			message += "Ошибка: невозможно получить температуру\n"
		} else {
			tempCelsius := convert.KelToCel(tempKelvin)
			message += fmt.Sprintf("Температура: %.f °C\n", tempCelsius)
		}

		rain, ok := firstItem["rain"]
		if ok && rain != nil {
			message += fmt.Sprintf("Дождь: %v\n", rain)
		} else {
			message += "Дождь: отсутствует\n"
		}

		snow, ok := firstItem["snow"]
		if ok && snow != nil {
			message += fmt.Sprintf("Снег: %v\n", snow)
		} else {
			message += "Снег: отсутствует\n"
		}

		windData, ok := firstItem["wind"].(map[string]any)
		if !ok {
			message += "Ошибка: невозможно получить данные о ветре\n"
			return message
		}

		windSpeed, ok := windData["speed"].(float64)
		if !ok {
			message += "Ошибка: невозможно получить скорость ветра\n"
		} else {
			message += fmt.Sprintf("Скорость ветра: %.2f м/c\n", windSpeed)
		}

		humidity, ok := mainData["humidity"].(float64)
		if !ok {
			message += "Ошибка: невозможно получить влажность\n"
		} else {
			message += fmt.Sprintf("Влажность: %.f %%\n", humidity)
		}

		cloudsData, ok := firstItem["clouds"].(map[string]any)
		if !ok {
			message += "Ошибка: невозможно получить данные об облачности\n"
			return message
		}
		clouds, ok := cloudsData["all"].(float64)
		if !ok {
			message += "Ошибка: невозможно получить облачность\n"
		} else {
			message += fmt.Sprintf("Облачность: %.f %%\n", clouds)
		}

		pressure, ok := mainData["pressure"].(float64)
		if !ok {
			message += "Ошибка: невозможно получить давление\n"
		} else {
			message += fmt.Sprintf("Давление: %.f гПа\n", pressure)
		}

	} else {
		message += "Ошибка: данные о погоде отсутствуют\n"
		return message
	}

	return message
}

func RandomInsult(data map[string]any) (message string) {
	message = data["insult"].(string)
	return message
}

func HowYouGosling() (message string) {
	message = strconv.Itoa(rand.Intn(100))
	return message
}

func ExchangeRates(data map[string]any) (message string) {

	listTicker := [...]string{"USD", "EUR", "CNY", "CHF", "JPY", "BYN", "HKD", "KZT", "UAH", "RSD", "PLN", "GBP"}
	time := fmt.Sprintf("Курс валют на %s\n\n", data["Date"].(string)[:10])
	message += time
	for _, ticker := range listTicker {
		valute := data["Valute"].(map[string]any)[ticker].(map[string]any)
		rate := valute["Value"].(float64) / valute["Nominal"].(float64)
		message += fmt.Sprintf("%s -> %.2f\n", valute["Name"], rate)

	}
	return message
}
func Game(bot *tgbotapi.BotAPI, update tgbotapi.Update, moveUser uint8) (message string) {
	var moveBot string
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ход бота")
	//отправляем сообщение
	SendMessage(bot, msg)

	randomNum := rand.Intn(3)
	switch randomNum {
	case 0:
		moveBot = "🪨"
	case 1:
		moveBot = "✂️"
	case 2:
		moveBot = "📄"
	}

	msg = tgbotapi.NewMessage(update.Message.Chat.ID, moveBot)
	//отправляем сообщение
	SendMessage(bot, msg)

	if uint8(randomNum) == moveUser {
		message = "Ничья"
	} else if (moveUser+1)%3 == uint8(randomNum) {
		message = "Победа твоя!"
	} else {
		message = "Я победил! хахааха"
	}

	return message
}
func SendMessage(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {
	if _, err := bot.Send(msg); err != nil {
		log.Printf("Ошибка при отправке сообщения: %v", err)
	}
}

// функция для отправки в группу кто какое сообщение отправилл боту
func SendGroupMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	groupMsg := tgbotapi.NewMessage(-1002065132697, fmt.Sprintf("%s \nОт: @%s", update.Message.Text, update.Message.From.UserName))
	if _, err := bot.Send(groupMsg); err != nil {
		log.Printf("Ошибка при отправке сообщения в группу: %v", err)
	}
}
