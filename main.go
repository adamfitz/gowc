package main

import (
	"fmt"
	"time"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("World Clock")

	// Simulating multiple timezones
	timezones := []string{"UTC", "Asia/Kolkata", "Europe/Berlin", "America/New_York", "Australia/Sydney"}

	// Create a container for the clocks horizontally
	clockContainer := container.NewHBox() // Horizontal container

	
	// Add each timezone's label and time
	for _, tz := range timezones {
		locLabel := widget.NewLabel(tz)  // Location name (e.g., "UTC")
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

		// Create a container for each timezone and time
		clockBox := container.NewVBox(centeredLocLabel, centeredTimeLabel) // Vertical box for location + time
		clockContainer.Add(clockBox) // Add each clock box to the horizontal container
	}

	// Show the window with the clock layout
	myWindow.SetContent(clockContainer)
	myWindow.ShowAndRun()
}
