package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os/exec"
)

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
