package main

import (
	// "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
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
	sidebar := my_widgets.MySideBar()
	toolbar := my_widgets.MyToolBar()

	// Add the widgets to the window
	main_grid := container.New(layout.NewGridLayout(7), sidebar)
	boarder_items := container.NewBorder(toolbar, nil, nil, nil, main_grid)
	myWindow.SetContent(boarder_items)
	myWindow.ShowAndRun()
}
