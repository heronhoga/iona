package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/heronhoga/iona/internal/hardware/usb"
)

func NewUSBList(manager usb.Manager) fyne.CanvasObject {
	devices, err := manager.List()
	if err != nil {
		return widget.NewLabel(err.Error())
	}

	list := widget.NewList(
		func() int {
			return len(devices)
		},
		
		func() fyne.CanvasObject {
			return container.NewVBox(
				widget.NewLabel("Product"),
				widget.NewLabel("Manufacturer"),
				widget.NewLabel("Status"),
			)
		},

		func(id widget.ListItemID, item fyne.CanvasObject) {
			device := devices[id]
			row := item.(*fyne.Container)

			product := row.Objects[0].(*widget.Label)
			manufacturer := row.Objects[1].(*widget.Label)
			status := row.Objects[2].(*widget.Label)

			product.SetText(device.Product)
			manufacturer.SetText(device.Manufacturer)

			if device.Authorized {
				status.SetText("Status: Enabled")
			} else {
				status.SetText("Status: Disabled")
			}
		},
	)

	return list
}