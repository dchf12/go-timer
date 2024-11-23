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

func main() {
	app := app.New()
	window := app.NewWindow("timer")
	window.Resize(fyne.NewSize(400, 400))

	entry := widget.NewEntry()
	entry.SetPlaceHolder("Enter time in seconds")

	startButton := widget.NewButton("Start", func() {
		seconds, err := strconv.Atoi(entry.Text)
		if err != nil {
			popup := widget.NewPopUp(canvas.NewText("秒数は整数で指定してください", nil), window.Canvas())
			popup.Show()
			time.AfterFunc(1*time.Second, func() {
				popup.Hide()
			})
			return
		}

		popup := widget.NewPopUp(canvas.NewText(fmt.Sprintf("%d秒のタイマー開始", seconds), nil), window.Canvas())
		popup.Show()
		time.AfterFunc(1*time.Second, func() {
			popup.Hide()
		})

		go func() {
			time.Sleep(time.Duration(seconds) * time.Second)
			completedPopup := widget.NewPopUp(canvas.NewText("タイマー終了", nil), window.Canvas())
			completedPopup.Show()
			time.AfterFunc(1*time.Second, func() {
				completedPopup.Hide()
			})
			_ = exec.Command("afplay", "/System/Library/Sounds/Ping.aiff").Run()
		}()
	})

	window.SetContent(container.NewVBox(
		widget.NewLabel("Timer"),
		entry,
		startButton,
	))

	window.ShowAndRun()
}
