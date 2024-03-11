package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func MyToolBar () *widget.Toolbar {
	sidebar_icon, err := fyne.LoadResourceFromPath("assets/sidebar-icon.png")
	if err != nil {
		panic("The icon filed to load: " + err.Error())
	}

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(sidebar_icon, helloWorld),
	)

	return toolbar
}