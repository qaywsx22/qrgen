package panels

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
)

type iconForQRImage struct {
}

func (d *iconForQRImage) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(MinQrIconSize, MinQrIconSize)
}

func (d *iconForQRImage) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	pos := fyne.NewPos(0, 0)
	size := d.MinSize(nil)
	for _, o := range objects {
		o.Resize(size)
		o.Move(pos)
	}
}

func PreviewPanel() fyne.CanvasObject {
	qrImg := canvas.NewImageFromFile(QrIconString)
	qrImg.FillMode = canvas.ImageFillStretch

	qrIcon := fyne.NewContainerWithLayout(&iconForQRImage{}, qrImg)
	qrPane := fyne.NewContainerWithLayout(layout.NewCenterLayout(), qrIcon)

	return qrPane
}
