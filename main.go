package main

import (
	// "fyne.io/fyne/v2"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"

	// "fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"

	// "fyne.io/fyne/v2/widget"
	my_widgets "docker_manager/widgets"
)

func main() {
	// Create the app
	myApp := app.New()
	myWindow := myApp.NewWindow("Docker Manager")

	// Create the widgets
	var sidebar *fyne.Container = my_widgets.MySideBar()
	var toolbar *widget.Toolbar = my_widgets.MyToolBar(sidebar)

	// Add the widgets to the window
	main_grid := container.New(layout.NewGridLayout(7), sidebar)
	boarder_items := container.NewBorder(toolbar, nil, nil, nil, main_grid)
	myWindow.SetContent(boarder_items)
	myWindow.ShowAndRun()
}
