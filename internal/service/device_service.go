package service

import (
	"github.com/heronhoga/iona/internal/device"
	"github.com/heronhoga/iona/internal/model"
)

type DeviceService struct {
    manager device.Manager
}

func NewDeviceService(manager device.Manager) *DeviceService {
	return &DeviceService{
		manager: manager,
	}
}

func (s *DeviceService) ListDevices() ([]model.Device, error) {
	return s.manager.ListDevices()
}

func (s *DeviceService) Enable(id string) error {
	return s.manager.Enable(id)
}

func (s *DeviceService) Disable(id string) error {
	return s.manager.Disable(id)
}