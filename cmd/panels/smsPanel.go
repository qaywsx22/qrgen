package panels

import (
	"fyne.io/fyne/widget"
)

type SMSPanel struct {
	widget.Form
	entryNr  *widget.Entry
	entryTxt *widget.Entry
}

func NewSMSPanel() *SMSPanel {
	sp := &SMSPanel{}
	enr := widget.NewEntry()
	etxt := widget.NewMultiLineEntry()

	enr.SetPlaceHolder("+1123456")
	etxt.SetPlaceHolder("Hello!")

	sp.Append(SmsNrLabelString, enr)
	sp.Append(SmsTextLabelString, etxt)
	sp.ExtendBaseWidget(sp)

	sp.entryNr = enr
	sp.entryTxt = etxt

	return sp
}

func (sp *SMSPanel) GetData() string {
	nr := sp.entryNr.Text
	txt := sp.entryTxt.Text
	return "SMSTO:" + nr + ":" + txt
}
