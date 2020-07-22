package panels

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func SMSPanel() fyne.CanvasObject {
	f := widget.NewForm()
	f.Append(SmsNrLabelString, widget.NewEntry())
	f.Append(SmsTextLabelString, widget.NewMultiLineEntry())

	return f
}
