package main

import (
	"encoding/base64"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("b64")

	txt := binding.NewString()
	entry := widget.NewMultiLineEntry()
	entry.SetMinRowsVisible(20)
	entry.Bind(txt)
	b := widget.NewButton("Encode/Decode", encodeDecode(txt))
	c := container.NewVBox(entry, b)
	w.SetContent(c)

	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}

func encodeDecode(txt binding.String) func() {
	return func() {
		str, err := txt.Get()
		if err != nil {
			return
		}

		output, err := base64.StdEncoding.DecodeString(str)
		if err == nil {
			txt.Set(string(output))
			return
		}

		txt.Set(base64.StdEncoding.EncodeToString([]byte(str)))
	}
}
