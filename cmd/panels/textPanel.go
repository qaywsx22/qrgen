package panels

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func TextPanel() fyne.CanvasObject {
	f := widget.NewForm()
	f.Append(TextLabelString, widget.NewMultiLineEntry())

	return f
}
