package usb

type USBDevice struct {
	ID           string
	Product      string
	Manufacturer string
	VendorID     string
	ProductID    string
	Authorized   bool
}