package utils

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type Device struct {
	Name string
	ID   string
}

func GetDevices() ([]Device, error) {
	// TODO get devices
	var devices []Device

	cmd := exec.Command("flutter", "devices")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return devices, err
	}

	lines := strings.Split(string(output), "\n")

	for i, line := range lines {
		if strings.Contains(line, "connected device") && i == 0 {
			devicecount := strings.Split(line, " ")[1]
			_, err := strconv.Atoi(devicecount)
			if err != nil {
				return devices, err
			}
			// if count == 1 {
			// 	fmt.Printf("Found %s device\n", devicecount)
			// } else {
			// 	fmt.Printf("Found %s devices\n", devicecount)
			// }
			continue
		}
		// Trim devices
		parts := strings.Split(line, "•")
		if len(parts) < 3 {
			continue
		}
		device := Device{
			Name: strings.TrimSpace(parts[0]),
			ID:   strings.TrimSpace(parts[1]),
		}

		devices = append(devices, device)
	}
	return devices, nil
}
