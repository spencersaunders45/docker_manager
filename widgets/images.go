package widgets

import (
	"fmt"
	"os"
	"os/exec"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	// "fyne.io/fyne/v2/widget"
)

func DockerImageList() *fyne.Container {
	var cmd *exec.Cmd = exec.Command("docker", "images")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("An error occurred while checking for images:, ", err)
		os.Exit(1)
	}
	fmt.Println(string(output))

	var image_window *fyne.Container = container.New(layout.NewGridLayout(3))
	return image_window
}
