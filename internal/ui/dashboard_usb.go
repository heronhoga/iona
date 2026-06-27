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

	refresh := widget.NewButton("Refresh", func() {
		if err := d.Refresh(); err != nil {
			dialog.ShowError(err, fyne.CurrentApp().Driver().AllWindows()[0])
		}
	})

	d.list = widget.NewList(
		func() int {
			return len(d.devices)
		},
		func() fyne.CanvasObject {
			product := widget.NewLabel("")
			manufacturer := widget.NewLabel("")
			authorized := widget.NewLabel("")

			enableBtn := widget.NewButton("Enable", nil)
			disableBtn := widget.NewButton("Disable", nil)

			return container.NewVBox(
				product,
				manufacturer,
				authorized,
				enableBtn,
				disableBtn,
			)
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			row := item.(*fyne.Container)

			product := row.Objects[0].(*widget.Label)
			manufacturer := row.Objects[1].(*widget.Label)
			authorized := row.Objects[2].(*widget.Label)
			enableBtn := row.Objects[3].(*widget.Button)
			disableBtn := row.Objects[4].(*widget.Button)

			device := d.devices[id]

			product.SetText("USB Product: " + device.Product)
			manufacturer.SetText("Manufacturer: " + device.Manufacturer)

			if device.Authorized {
				authorized.SetText("Authorized: YES")

				enableBtn.Hide()
				disableBtn.Show()

				disableBtn.OnTapped = func() {
					if err := d.manager.Disable(device.ID); err != nil {
						dialog.ShowError(err, fyne.CurrentApp().Driver().AllWindows()[0])
						return
					}
					d.Refresh()
				}
			} else {
				authorized.SetText("Authorized: NO")

				disableBtn.Hide()
				enableBtn.Show()

				enableBtn.OnTapped = func() {
					if err := d.manager.Enable(device.ID); err != nil {
						dialog.ShowError(err, fyne.CurrentApp().Driver().AllWindows()[0])
						return
					}
					d.Refresh()
				}
			}
		},
)

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