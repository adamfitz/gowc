package main

import (
	"fmt"
	"image/color"
	"time"
	//"fyne.io/fyne/v2"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("World Clock")

	// center window on startup
	myWindow.CenterOnScreen()

	// Disable resizing the window
	myWindow.SetFixedSize(true)

	// Simulating multiple timezones
	timezones := []string{"UTC", "Asia/Kolkata", "Europe/Berlin", "America/New_York", "Australia/Sydney"}

	colours := []color.Color{
		color.RGBA{0, 255, 0, 255},   // Green
		color.RGBA{0, 0, 255, 255},   // Blue
		color.RGBA{255, 0, 0, 255},   // Red
		color.RGBA{255, 255, 0, 255}, // Yellow
		color.RGBA{0, 255, 255, 255}, // Cyan
	}

	// Create a container for the clocks horizontally
	clockContainer := container.NewHBox() // Horizontal container

	// Create a container for the clocks without layout for precise control
	//clockContainer := container.NewWithoutLayout()

	// Add each timezone's label and time
	for i, tz := range timezones {
		locLabel := widget.NewLabel(tz)       // Location name (e.g., "UTC")
		timeLabel := widget.NewLabel("00:00") // Placeholder time

		// Location label// Create a variable to store the location label
		var locationLabel = ""
		// Case to translate the timezone name I want to see
		switch tz {
		case "UTC":
			locationLabel = "UTC"
		case "Asia/Kolkata":
			locationLabel = "BLR"
		case "Europe/Berlin":
			locationLabel = "ERD"
		case "America/New_York":
			locationLabel = "MIA"
		case "Australia/Sydney":
			locationLabel = "LT"
		}
		// Update the location label text
		locLabel.SetText(locationLabel)

		// Center the location label
		centeredLocLabel := container.NewCenter(locLabel)

		// Update the time inside a go routine or update loop
		go func(tz string, timeLabel *widget.Label) {
			for {
				locTime, err := time.LoadLocation(tz)
				if err != nil {
					fmt.Println("Error loading timezone:", err)
					continue
				}
				currentTime := time.Now().In(locTime).Format("15:04:05")
				timeLabel.SetText(currentTime)
				time.Sleep(time.Second) // Update every second
			}
		}(tz, timeLabel)

		// Center the time label
		centeredTimeLabel := container.NewCenter(timeLabel)

		// Create a container for each timezone and time with a colored background
		rect := canvas.NewRectangle(colours[i%len(colours)]) // Use modulo to cycle through colors
		clockBox := container.NewVBox(centeredLocLabel, centeredTimeLabel)
		clockWithBackground := container.NewStack(rect, clockBox) // Add background rectangle

		clockContainer.Add(clockWithBackground)

	}

	// Show the window with the clock layout
	myWindow.SetContent(clockContainer)
	myWindow.ShowAndRun()
}
