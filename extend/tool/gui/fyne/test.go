// developer.fyne.io
package main

import (
	"fyne.io/fyne/v2"
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Choice Widgets")

	check := widget.NewCheckGroup([]string{"check_1", "check_2"}, func(value []string) {
		log.Println("Check set to", value)
	})
	radio := widget.NewRadioGroup([]string{"Option 1", "Option 2"}, func(value string) {
		log.Println("Radio set to", value)
	})
	combo := widget.NewSelect([]string{"Option 1", "Option 2"}, func(value string) {
		log.Println("Select set to", value)
	})

	myWindow.SetContent(container.NewVBox(check, radio, combo))
	myWindow.Resize(fyne.NewSize(300, 500))
	myWindow.ShowAndRun()
}
