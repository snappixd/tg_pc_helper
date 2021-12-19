package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os/exec"
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

func powerOff(msg *tgbotapi.MessageConfig) {
	msg.Text = "Good Bye *) xd"

	cmd := exec.Command("cmd.exe", "/C", "shutdown /t 0", "/s")

	if err := cmd.Run(); err != nil {
		fmt.Println("Failed to initiate shutdown:", err)
	}
}

func sleep(msg *tgbotapi.MessageConfig) {
	msg.Text = "Good night"

	cmd := exec.Command("cmd.exe", "/C", "shutdown /h")

	if err := cmd.Run(); err != nil {
		fmt.Println("Failed to sleep", err)
	}
}

func closeWindows(msg *tgbotapi.MessageConfig) {
	msg.Text = "All windows were closed, and in ur room too, xd"

}

func reboot(msg *tgbotapi.MessageConfig) {
	msg.Text = "Rebooting..."

	cmd := exec.Command("cmd.exe", "/C", "shutdown /r /t 0")

	if err := cmd.Run(); err != nil {
		fmt.Println("Failed to initiate shutdown:", err)
	}
}

func volumeUp(msg *tgbotapi.MessageConfig) {
	msg.Text = "+5%"
	cmd := exec.Command("cmd.exe", "/C", "nircmd.exe changesysvolume 3300")

	if err := cmd.Run(); err != nil {
		fmt.Println("Failed to change volume", err)
	}
}

func volumeDown(msg *tgbotapi.MessageConfig) {
	msg.Text = "-5%"

	cmd := exec.Command("cmd.exe", "/C", "nircmd.exe changesysvolume -3300")

	if err := cmd.Run(); err != nil {
		fmt.Println("Failed to change volume", err)
	}
}

func volumeMute(msg *tgbotapi.MessageConfig) {
	msg.Text = "Volume muted/unmuted"

	cmd := exec.Command("cmd.exe", "/C", "nircmd.exe mutesysvolume 2")

	if err := cmd.Run(); err != nil {
		fmt.Println("Failed to mute volume", err)
	}
}

func monitorOff(msg *tgbotapi.MessageConfig) {
	msg.Text = "Monitor offed"

	cmd := exec.Command("cmd.exe", "/C", "nircmd.exe monitor off")

	if err := cmd.Run(); err != nil {
		fmt.Println("Failed to off monitor", err)
	}
}
