package widgets

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func MySideBar() *fyne.Container {
	images_dropdown := widget.NewButton("Images", helloWorld)
	containers_dropdown := widget.NewButton("Containers", helloWorld)
	networks_dropdown := widget.NewButton("Networks", helloWorld)
	volumes_dropdown := widget.NewButton("Volumes", helloWorld)
	sidebar := container.NewVBox(images_dropdown, containers_dropdown, networks_dropdown, volumes_dropdown)
	
	return sidebar
}

func helloWorld() {
	fmt.Println("Hello World")
}