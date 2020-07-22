package panels

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func customizeQR() {
}

func showURLPanel() {
	ShowDataTypeForm(DTURL)
}

func showTextPanel() {
	ShowDataTypeForm(DTText)
}

func showSMSPanel() {
	ShowDataTypeForm(DTSMS)
}

func NavigatorPanel() fyne.CanvasObject {
	// btnColor := widget.NewButton(AdjustColorsString, adjustColors)
	// btnDesign := widget.NewButton(AdjustDesignString, adjustDesign)
	btnCustomize := widget.NewButton(CustomizeQRCodeString, customizeQR)

	ac := widget.NewGroup(DTContainerTitleString)
	icon, _ := fyne.LoadResourceFromPath("assets/url_1.png")
	ac.Append(widget.NewButtonWithIcon(DTURLString, icon, showURLPanel))
	icon, _ = fyne.LoadResourceFromPath("assets/text_2.png")
	ac.Append(widget.NewButtonWithIcon(DTTextString, icon, showTextPanel))
	icon, _ = fyne.LoadResourceFromPath("assets/message_three_points.png")
	ac.Append(widget.NewButtonWithIcon(DTSMSString, icon, showSMSPanel))

	np := widget.NewVBox()
	// np.Append(btnColor)
	// np.Append(btnDesign)
	np.Append(btnCustomize)
	np.Append(ac)

	return np
}
