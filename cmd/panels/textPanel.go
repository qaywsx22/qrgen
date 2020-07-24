package panels

import (
	"fyne.io/fyne/widget"
)

type TextPanel struct {
	widget.Form
	entry *widget.Entry
}

func NewTextPanel() *TextPanel {
	e := widget.NewMultiLineEntry()
	e.SetPlaceHolder("Enter your text here\ncan be multiline")

	tp := &TextPanel{}
	tp.ExtendBaseWidget(tp)
	tp.Append(TextLabelString, e)
	tp.entry = e

	return tp
}

// const (
// 	testString = "one two три четыре fünf sechs"
// )

func (tp *TextPanel) GetData() string {
	return tp.entry.Text
	// return testString
}
