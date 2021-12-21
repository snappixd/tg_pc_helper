package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"power_off_bot/cfg"
)

var keyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Volume +5%"),
		tgbotapi.NewKeyboardButton("Volume -5%"),
		tgbotapi.NewKeyboardButton("Mute/Unmute"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Off"),
		tgbotapi.NewKeyboardButton("Reboot"),
		tgbotapi.NewKeyboardButton("Sleep"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Monitor off"),
		tgbotapi.NewKeyboardButton("Close windows"),
	),
)


func main() {
	bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	checkUpdates(updates, bot)
}

func checkUpdates(updates tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI) {
	for update := range updates {
		if update.Message == nil { // If we got a message
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Поки що приймаю тільки команди!\nНажимай кнопочки *)")

		switch update.Message.Command() {
		case "start":
			msg.ReplyMarkup = keyboard
		}

		switch update.Message.Text {
		case "Sleep":
			sleep(&msg)
		case "Reboot":
			reboot(&msg)
		case "Off":
			powerOff(&msg)
		case "Close windows":
			closeWindows(&msg)
		case "Volume +5%":
			volumeUp(&msg)
		case "Volume -5%":
			volumeDown(&msg)
		case "Mute/Unmute":
			volumeMute(&msg)
		case "Monitor off":
			monitorOff(&msg)
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}


	}
}

