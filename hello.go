package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	myApp := app.New()

	window := myApp.NewWindow("Hello")

	window.SetContent(widget.NewVBox(
		widget.NewLabel("Hello test!"),
		widget.NewButton("Exit", func() {
			myApp.Quit()
		}),
	))

	window.ShowAndRun()
}
