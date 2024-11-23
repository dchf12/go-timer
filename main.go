package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const (
	Second = "秒"
	Minute = "分"
)

func main() {
	app := app.New()
	window := app.NewWindow("timer")
	window.Resize(fyne.NewSize(400, 400))

	entry := widget.NewEntry()
	entry.SetPlaceHolder("Enter time in seconds")

	secondsButton := widget.NewButton("秒で開始", func() {
		handleTimer(entry.Text, Second, window)
	})

	minutesButton := widget.NewButton("分で開始", func() {
		handleTimer(entry.Text, Minute, window)
	})

	window.SetContent(container.NewVBox(
		widget.NewLabel("Timer"),
		entry,
		container.NewHBox(
			secondsButton,
			minutesButton,
		),
	))

	window.ShowAndRun()
}

func handleTimer(input string, unit string, window fyne.Window) {
	timeValue, err := strconv.Atoi(input)
	if err != nil {
		showPopup(window, fmt.Sprintf("%sは整数で指定してください", unit))
		return
	}

	showPopup(window, fmt.Sprintf("%d%sのタイマー開始", timeValue, unit))

	go func() {
		var duration time.Duration
		if unit == Second {
			duration = time.Duration(timeValue) * time.Second
		} else if unit == Minute {
			duration = time.Duration(timeValue) * time.Minute
		}
		time.Sleep(duration)
		showPopup(window, "タイマー終了")
		playCompletionSound()
	}()
}

func showPopup(window fyne.Window, message string) {
	popup := widget.NewPopUp(canvas.NewText(message, nil), window.Canvas())
	popup.Show()
	time.AfterFunc(1*time.Second, func() {
		popup.Hide()
	})
}

func playCompletionSound() {
	_ = exec.Command("afplay", "/System/Library/Sounds/Ping.aiff").Run()
}
