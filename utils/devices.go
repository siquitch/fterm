package utils

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

type Device struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func GetDevices() ([]Device, error) {
	var devices []Device

	cmd := exec.Command("flutter", "devices", "--machine")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return devices, err
	}

	err = json.Unmarshal(output, &devices)

	if err != nil {
		return devices, err
	}

	return devices, nil
}

func GetEmulators() ([]Device, error) {
	var devices []Device

	cmd := exec.Command("flutter", "emulators")

	output, err := cmd.Output()

	if err != nil {
		return devices, err
	}

	lines := strings.Split(string(output), "\n")

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

	return devices, nil
}
