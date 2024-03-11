package widgets

import (
	// "fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func MyToolBar (sidebar *fyne.Container) *widget.Toolbar {
	// Load in icons
	sidebar_icon, err := fyne.LoadResourceFromPath("assets/sidebar-icon.png")
	if err != nil {
		panic("The icon filed to load: " + err.Error())
	}

	// Create the toolbar
	var toolbar *widget.Toolbar = widget.NewToolbar(
		widget.NewToolbarAction(sidebar_icon, func() {
			toggleSideBar(sidebar)
		}),
	)

	return toolbar
}

/* Toggles the visibility of the sidebar

*/
func toggleSideBar(sidebar *fyne.Container) {
	if sidebar.Visible() {
		sidebar.Hide()
	} else {
		sidebar.Show()
	}
}