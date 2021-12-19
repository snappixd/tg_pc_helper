package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os/exec"
)

var keyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Volume +5%"),
		tgbotapi.NewKeyboardButton("Volume -5%"),
		tgbotapi.NewKeyboardButton("Off"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Reboot"),
		tgbotapi.NewKeyboardButton("Sleep"),
		tgbotapi.NewKeyboardButton("Close windows"),
	),
)


func main() {
	bot, err := tgbotapi.NewBotAPI("5087575563:AAH997vLcthQJB3pE3fDRnorl1OhU_DS5Ls")
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

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		switch update.Message.Command() {
		case "start":
			msg.ReplyMarkup = keyboard
		case "help":
			msg.Text = "Є кілька команд: \n/help - наявні команди\n/off - офнути пк\n/sleep - в сон\n/reboot - перезагрузка\n/volume_up - звук +5%\n/volume_down - звук -5%"
		}

		switch update.Message.Text {
		case "Sleep":
			sleep(msg)
		case "Reboot":
			reboot(msg)
		case "Off":
			powerOff(msg)
		case "Close windows":
			closeWindows(msg)
		case "Volume +5%":
			volumeUp(msg)
		case "Volume -5%":
			volumeDown(msg)
		}


		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}

func powerOff(msg tgbotapi.MessageConfig) {
	msg.Text = "Good Bye *) xd"

	cmd := exec.Command("cmd.exe", "/C", "shutdown /t 0", "/s")

	if err := cmd.Run(); err != nil {
		fmt.Println("Failed to initiate shutdown:", err)
	}
}

func sleep(msg tgbotapi.MessageConfig) {
	msg.Text = "Good night"

	cmd := exec.Command("cmd.exe", "/C", "shutdown /h")

	if err := cmd.Run(); err != nil {
		fmt.Println("Failed to sleep", err)
	}
}

func closeWindows(msg tgbotapi.MessageConfig) {
	msg.Text = "All windows were closed, and in ur room too, xd"

}

func reboot(msg tgbotapi.MessageConfig) {
	msg.Text = "Rebooting..."

	cmd := exec.Command("cmd.exe", "/C", "shutdown /r /t 0")

	if err := cmd.Run(); err != nil {
		fmt.Println("Failed to initiate shutdown:", err)
	}
}

func volumeUp(msg tgbotapi.MessageConfig) {
	msg.Text = "Volume +5%"

	cmd := exec.Command("cmd.exe", "/C", "nircmd.exe changesysvolume 3300")

	if err := cmd.Run(); err != nil {
		fmt.Println("Failed to change volume", err)
	}
}

func volumeDown(msg tgbotapi.MessageConfig) {
	msg.Text = "Volume -5%"

	cmd := exec.Command("cmd.exe", "/C", "nircmd.exe changesysvolume -3300")

	if err := cmd.Run(); err != nil {
		fmt.Println("Failed to change volume", err)
	}
}
