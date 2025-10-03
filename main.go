package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Create a new application
	myApp := app.New()
	myWindow := myApp.NewWindow("Fyne Bug Demo")
	myWindow.Resize(fyne.NewSize(400, 300))

	// Create the main content - Hello World message
	helloLabel := widget.NewLabel("Hello, World!")
	helloLabel.Alignment = fyne.TextAlignCenter

	// Create the main content container
	content := container.NewCenter(helloLabel)

	// Create File menu
	fileMenu := fyne.NewMenu("File",
		fyne.NewMenuItem("New", func() {
			// Placeholder action
		}),
		fyne.NewMenuItem("Open", func() {
			// Placeholder action
		}),
		fyne.NewMenuItem("Save", func() {
			// Placeholder action
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Exit", func() {
			myApp.Quit()
		}),
	)

	// Create Edit menu
	editMenu := fyne.NewMenu("Edit",
		fyne.NewMenuItem("Cut", func() {
			// Placeholder action
		}),
		fyne.NewMenuItem("Copy", func() {
			// Placeholder action
		}),
		fyne.NewMenuItem("Paste", func() {
			// Placeholder action
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Select All", func() {
			// Placeholder action
		}),
	)

	// Create View menu
	viewMenu := fyne.NewMenu("View",
		fyne.NewMenuItem("Zoom In", func() {
			// Placeholder action
		}),
		fyne.NewMenuItem("Zoom Out", func() {
			// Placeholder action
		}),
		fyne.NewMenuItem("Reset Zoom", func() {
			// Placeholder action
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Full Screen", func() {
			// Placeholder action
		}),
	)

	// Create the main menu
	mainMenu := fyne.NewMainMenu(fileMenu, editMenu, viewMenu)

	// Set the menu and content
	myWindow.SetMainMenu(mainMenu)
	myWindow.SetContent(content)

	// Show the window and run the app
	myWindow.ShowAndRun()
}
