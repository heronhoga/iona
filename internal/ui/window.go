package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/heronhoga/iona/internal/hardware/usb"
)

func Run() {
	app := app.New()

	w := app.NewWindow("IONA @1afrinata")
	w.Resize(fyne.NewSize(1000, 800))
	w.CenterOnScreen()

	manager := usb.NewLinuxManager()
	dashboard := NewUSBDashboard(manager)
	w.SetContent(dashboard.USBView())
	w.ShowAndRun()
}
