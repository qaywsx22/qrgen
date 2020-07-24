package panels

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/qaywsx22/qr/cmd/logic"
)

type DataProvider interface {
	GetData() string
}

var (
	urlPanel     fyne.CanvasObject
	textPanel    fyne.CanvasObject
	smsPanel     fyne.CanvasObject
	urlIcon      fyne.Resource
	textIcon     fyne.Resource
	smsIcon      fyne.Resource
	dataTypeIcon *widget.Icon
)

func DataPanel() fyne.CanvasObject {
	urlPanel = NewURLPanel()
	urlIcon, _ = fyne.LoadResourceFromPath("assets/url_1.png")

	textPanel = NewTextPanel()
	textIcon, _ = fyne.LoadResourceFromPath("assets/text_2.png")

	smsPanel = NewSMSPanel()
	smsIcon, _ = fyne.LoadResourceFromPath("assets/message_three_points.png")

	dataTypeIcon = widget.NewIcon(urlIcon)
	dataTypeIconPanel := fyne.NewContainerWithLayout(layout.NewCenterLayout(), dataTypeIcon)

	dataPane := widget.NewGroup(DataPaneTitlelString)
	dataPane.Append(urlPanel)
	dataPane.Append(textPanel)
	dataPane.Append(smsPanel)

	dp := widget.NewVBox(dataTypeIconPanel, dataPane)
	hideAll()

	return dp
}

func hideAll() {
	urlPanel.Hide()
	textPanel.Hide()
	smsPanel.Hide()
}

func ShowDataTypeForm(dtype int) {
	hideAll()
	switch dtype {
	case DTText:
		textPanel.Show()
		dataTypeIcon.SetResource(textIcon)
	case DTURL:
		urlPanel.Show()
		dataTypeIcon.SetResource(urlIcon)
	case DTSMS:
		smsPanel.Show()
		dataTypeIcon.SetResource(smsIcon)
	}
}

func Generate() {
	var s string = "Not initialized"
	if urlPanel.Visible() {
		s = urlPanel.(DataProvider).GetData()
	} else if textPanel.Visible() {
		s = textPanel.(DataProvider).GetData()
	} else if smsPanel.Visible() {
		s = smsPanel.(DataProvider).GetData()
	}
	logic.GenerateQRCode(s)
}
