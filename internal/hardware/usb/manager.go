package usb

type Manager interface {
	List() ([]USBDevice, error)
	Enable(id string) error
	Disable(id string) error
}
