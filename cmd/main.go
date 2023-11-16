package main

import (
	"log"
	"os"
	"weather-bot/pkg/telegram"

	"github.com/briandowns/openweathermap"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		log.Fatalln(err)
	}

	bot.Debug = true

	w, err := openweathermap.NewCurrent("F", "ru", os.Getenv("OPENWEATHER_TOKEN"))
	if err != nil {
		log.Fatalln(err)
	}

	w.CurrentByName("Grozny")

	telegramBot := telegram.NewBot(bot, w)

	if err := telegramBot.Start(); err != nil {
		log.Fatalln(err)
	}

}
