package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/qaywsx22/qr/cmd/panels"
)

func main() {
	app := app.New()
	app.Settings().SetTheme(theme.LightTheme())
	w := app.NewWindow(panels.AppTitle)
	mainPane := widget.NewVBox(
		panels.PreviewPanel(),
		panels.DataPanel(),
	)
	tb := panels.CreateToolbar(w)
	navPane := panels.NavigatorPanel()
	contentPane := fyne.NewContainerWithLayout(
		layout.NewBorderLayout(tb, nil, navPane, nil),
		tb,
		navPane,
		mainPane,
	)
	w.SetContent(contentPane)
	wSize := fyne.NewSize(panels.MainWindowWidth, panels.MainWindowHeight)
	w.Resize(wSize)

	panels.Init()

	w.ShowAndRun()
}
