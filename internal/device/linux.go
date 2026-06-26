package device

import (
	"bytes"
	"os/exec"
	"regexp"

	"github.com/heronhoga/iona/internal/model"
)

type LinuxManager struct{}

func NewLinuxManager() Manager {
	return &LinuxManager{}
}

func (m *LinuxManager) ListDevices() ([]model.Device, error) {
    cmd := exec.Command("xinput", "list")
    var out bytes.Buffer
    cmd.Stdout = &out

    if err := cmd.Run(); err != nil {
    return nil, err
    }

    return parseDevices(out.String()), nil
}

func (m *LinuxManager) Enable(id string) error {
	cmd := exec.Command("xinput", "enable", id)
	return cmd.Run()
}

func (m *LinuxManager) Disable(id string) error {
	cmd := exec.Command("xinput", "disable", id)
	return cmd.Run()
}

func parseDevices(output string) []model.Device {
	re := regexp.MustCompile(`(.+?)\s+id=(\d+)`)

	var devices []model.Device

	for _, line := range regexp.MustCompile("\n").Split(output, -1) {
		match := re.FindStringSubmatch(line)

		if len(match) != 3 {
			continue
		}

		devices = append(devices, model.Device{
			Name: match[1],
			ID:   match[2],
		})
	}

	return devices
}