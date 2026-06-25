package device

import "github.com/heronhoga/iona/internal/model"

type Manager interface {
    ListDevices() ([]model.Device, error)
    Enable(id string) error
    Disable(id string) error
}