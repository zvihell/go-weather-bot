package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	commandStart = "start"
	weatherStart = "weather"
	replyStart   = "Привет!Введи команду /weather для получения информации о погоде в Грозном"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) {

	switch message.Command() {
	case commandStart:
		b.handleStartCommand(message)
	case weatherStart:
		b.weatherCommand(message)
	default:
		b.handleUnknownCommand(message)
	}

}

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)

	b.bot.Send(msg)

}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, replyStart)

	b.bot.Send(msg)

}

func (b *Bot) weatherCommand(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.openMap.Weather[0].Description)
	b.bot.Send(msg)

}
func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Я не знаю такой команды")

	b.bot.Send(msg)

}
