package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/heronhoga/iona/internal/hardware/usb"
)

func Run() {
	app := app.New()

	w := app.NewWindow("IONA @1afrinata")
	w.Resize(fyne.NewSize(600, 400))

	manager := usb.NewLinuxManager()
	dashboard := NewUSBDashboard(manager)
	w.SetContent(dashboard.USBView())
	w.ShowAndRun()
}
