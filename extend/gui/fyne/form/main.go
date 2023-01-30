package main

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Form Widget")

	// input
	entry := widget.NewEntry()
	entry.Validator = validation.NewRegexp(`^\d+$`, "Input is not a valid date")

	// textArea
	textArea := widget.NewMultiLineEntry()
	textArea.Validator = validation.NewAllStrings(func(s string) error {
		if len(s) > 20 {
			return errors.New("size over")
		}
		return nil
	})

	// select
	selector := widget.NewSelect([]string{"a", "b", "c"}, func(string) {})

	// checkbox
	check := widget.NewCheckGroup([]string{"Option A", "Option B", "Option C"}, nil)

	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "Entry", Widget: entry}},
		OnSubmit: func() { // optional, handle form submission
			log.Println("Form submitted:", entry.Text)
			log.Println("multiline:", textArea.Text)
			log.Println("selector:", selector.Selected)
			log.Println("check:", check.Selected)
			myWindow.Close()
		},
		OnCancel: func() {
			if err := entry.Validate(); err != nil {
				entry.SetText("")
			}
			textArea.SetText("test 666")
			selector.ClearSelected()
			check.SetSelected([]string{})
		},
	}

	// we can also append items
	form.Append("Text", textArea)
	form.Append("selector", selector)
	form.Append("check", check)

	myWindow.SetContent(form)
	myWindow.Resize(fyne.NewSize(300, 500))

	// 弹框
	d := dialog.NewInformation("test", "msg", myWindow)
	d.Resize(fyne.NewSize(300, 500))
	d.Show()

	myWindow.ShowAndRun()
}
