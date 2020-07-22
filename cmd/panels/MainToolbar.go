package panels

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"

	// "fyne.io/fyne/storage"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func fileSaved(f fyne.URIWriteCloser) {
	if f == nil {
		fmt.Println("Cancelled")
		return
	}

	fmt.Println("Save to...", f.URI())
}

func CreateToolbar(win fyne.Window) fyne.CanvasObject {
	tb := widget.NewToolbar()

	icon, err := fyne.LoadResourceFromPath("assets/shut_down.png")
	if err != nil {
		icon = theme.CancelIcon()
	}
	act := widget.NewToolbarAction(icon, func() {
		fyne.CurrentApp().Quit()
	})

	tb.Append(act)

	icon, err = fyne.LoadResourceFromPath("assets/save_2.png")
	act = widget.NewToolbarAction(icon, func() {
		fd := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}

			fileSaved(writer)
		}, win)
		// fd.SetFilter(storage.NewExtensionFileFilter([]string{".png", ".svg"}))
		fd.Show()
	})
	tb.Append(act)

	tb.Append(widget.NewToolbarSeparator())

	icon, err = fyne.LoadResourceFromPath("assets/arrow_repeat.png")
	act = widget.NewToolbarAction(icon, func() {
	})
	tb.Append(act)

	tb.Append(widget.NewToolbarSpacer())

	icon, err = fyne.LoadResourceFromPath("assets/settings_1.png")
	act = widget.NewToolbarAction(icon, func() {
	})
	tb.Append(act)

	return tb
}
