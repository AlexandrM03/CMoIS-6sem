package main

import (
	"crypto/sha256"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"time"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("SHA256 Tool")

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter message to be hashed")
	input.Resize(fyne.NewSize(200, 50))

	hashLabel := widget.NewLabel("")
	timeLabel := widget.NewLabel("")

	button := widget.NewButton("Hash", func() {
		message := []byte(input.Text)

		start := time.Now()
		hash := sha256.Sum256(message)
		elapsed := time.Since(start)

		hashLabel.SetText(fmt.Sprintf("Hash: %x", hash))
		timeLabel.SetText(fmt.Sprintf("Elapsed time: %v", elapsed))
	})

	content := container.New(layout.NewVBoxLayout(), input, button, hashLabel, timeLabel)
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(400, 400))
	myWindow.ShowAndRun()
}
