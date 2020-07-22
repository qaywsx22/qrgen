package panels

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func URLPanel() fyne.CanvasObject {
	f := widget.NewForm()
	f.Append(UrlLabelString, widget.NewEntry())

	return f
}
