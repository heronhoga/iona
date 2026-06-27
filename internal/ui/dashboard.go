package ui

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/container"
)

func NewDashboard(
    usb *USBDashboard,
) fyne.CanvasObject {

    tabs := container.NewAppTabs(
        container.NewTabItem("USB", usb.USBView()),
    )

    return tabs
}