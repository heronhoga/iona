package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/heronhoga/iona/internal/hardware/usb"
)

type USBDashboard struct {
	manager usb.Manager
	devices []usb.USBDevice
	list *widget.List
}

func NewUSBDashboard(manager usb.Manager) *USBDashboard {
	return &USBDashboard{
		manager: manager,
	}
}

func (d *USBDashboard) USBView() fyne.CanvasObject {
	if err := d.Refresh(); err != nil {
		return widget.NewLabel(err.Error())
	}

	d.list = widget.NewList(
		func() int {
			return len(d.devices)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(d.devices[id].Product)
		},
	)

	refresh := widget.NewButton("Refresh", func() {
		if err := d.Refresh(); err != nil {
			dialog.ShowError(err, fyne.CurrentApp().Driver().AllWindows()[0])
		}
	})

	return container.NewBorder(
		refresh,
		nil,
		nil,
		nil,
		d.list,
	)
}


func (d *USBDashboard) Refresh() error {
	devices, err := d.manager.List()
	if err != nil {
		return err
	}

	d.devices = devices

	if d.list != nil {
		d.list.Refresh()
	}

	return nil
}