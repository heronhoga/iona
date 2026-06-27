package usb

import (
	"os"
	"path/filepath"
	"strings"
)

type LinuxManager struct{}

func NewLinuxManager() Manager {
	return &LinuxManager{}
}

func (m *LinuxManager) List() ([]USBDevice, error) {
	entries, err := os.ReadDir("/sys/bus/usb/devices")
	if err != nil {
		return nil, err
	}

	var devices []USBDevice

	for _, entry := range entries {
		name := entry.Name()

		if strings.HasPrefix(name, "usb") {
			continue
		}

		if strings.Contains(name, ":") {
			continue
		}

		devicePath := filepath.Join("/sys/bus/usb/devices", name)

		device := USBDevice{
			ID:           name,
			Product:      readFile(filepath.Join(devicePath, "product")),
			Manufacturer: readFile(filepath.Join(devicePath, "manufacturer")),
			VendorID:     readFile(filepath.Join(devicePath, "idVendor")),
			ProductID:    readFile(filepath.Join(devicePath, "idProduct")),
			Authorized:   readFile(filepath.Join(devicePath, "authorized")) == "1",
		}

		devices = append(devices, device)
	}

	return devices, nil
}

func (m *LinuxManager) Enable(id string) error {
	return nil
}

func (m *LinuxManager) Disable(id string) error {
	return nil
}
