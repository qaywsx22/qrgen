package panels

import (
	"fyne.io/fyne/widget"
)

type URLPanel struct {
	widget.Form
	entry *widget.Entry
}

func NewURLPanel() *URLPanel {
	up := &URLPanel{}
	up.ExtendBaseWidget(up)

	e := widget.NewEntry()
	up.Append(UrlLabelString, e)
	e.SetPlaceHolder("https://google.com")

	up.entry = e

	return up
}

func (up *URLPanel) GetData() string {
	return up.entry.Text
}
