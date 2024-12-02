package utils

import (
	"encoding/json"
	"strings"
)

type Device struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func ParseDevices(bytes []byte) ([]Device, error) {
	var devices []Device
	err := json.Unmarshal(bytes, &devices)

	if err != nil {
		return devices, err
	}

	return devices, nil
}

// Returns a list of available devices to use
//
// Deprecated: Use ParseDevices
func GetDevices() ([]Device, error) {
	var devices []Device

	cmd := FlutterDevices()
	output, err := cmd.Output()
	if err != nil {
		PrintError(err.Error())
		return devices, err
	}

	err = json.Unmarshal(output, &devices)

	if err != nil {
		return devices, err
	}

	return devices, nil
}

func ParseEmulators(bytes []byte) ([]Device, error) {
	var devices []Device

	lines := strings.Split(string(bytes), "\n")

	for i, line := range lines {
		if line == "" {
			continue
		}
		// No useful info on these lines
		if i >= 0 && i < 3 {
			continue
		}

		// Emulators start on line 4

		if line == "" {
			break
		}

		parts := strings.Split(line, "•")

		if len(parts) < 4 {
			continue
		}

		device := Device{
			ID:   strings.TrimSpace(parts[0]),
			Name: strings.TrimSpace(parts[1]),
		}

		devices = append(devices, device)
	}

	// Remove the first element which is "Name"
	devices = devices[0:]

	return devices, nil
}
