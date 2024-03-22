package widgets

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	// "fyne.io/fyne/v2/widget"
)

func DockerImageList() *fyne.Container {
	var image_window *fyne.Container = container.New(layout.NewGridLayout(1))
	// Run command to get list of images
	var cmd *exec.Cmd = exec.Command("docker", "images")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("An error occurred while checking for images:, ", err)
		os.Exit(1)
	}
	// Add text output to labels
	var results string = string(output)
	lines := strings.Split(results, "\n")
	for index, line := range lines {
		if index == 0 {
			continue
		} else {
			image_window.Add(widget.NewLabel(line))
		}
		
	}

	
	return image_window
}
