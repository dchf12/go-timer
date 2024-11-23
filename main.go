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

	secondsButton := widget.NewButton("秒で開始", func() {
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

			// タイマー完了時に音を鳴らす
			_ = exec.Command("afplay", "/System/Library/Sounds/Ping.aiff").Run()
		}()
	})

	minutesButton := widget.NewButton("分で開始", func() {
		minutes, err := strconv.Atoi(entry.Text)
		if err != nil {
			popup := widget.NewPopUp(canvas.NewText("分数は整数で指定してください", nil), window.Canvas())
			popup.Show()
			time.AfterFunc(1*time.Second, func() {
				popup.Hide()
			})
			return
		}

		popup := widget.NewPopUp(canvas.NewText(fmt.Sprintf("%d分のタイマー開始", minutes), nil), window.Canvas())
		popup.Show()
		time.AfterFunc(1*time.Second, func() {
			popup.Hide()
		})

		go func() {
			time.Sleep(time.Duration(minutes) * time.Minute)
			completedPopup := widget.NewPopUp(canvas.NewText("タイマー終了", nil), window.Canvas())
			completedPopup.Show()
			time.AfterFunc(1*time.Second, func() {
				completedPopup.Hide()
			})

			// タイマー完了時に音を鳴らす
			_ = exec.Command("afplay", "/System/Library/Sounds/Ping.aiff").Run()
		}()
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
