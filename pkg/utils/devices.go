package utils

import (
	"encoding/json"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
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

func GetDeviceTable(devices []Device) table.Model {
	c := []table.Column{
		{Title: "Name", Width: 20},
		{Title: "ID", Width: 20},
	}

	var r []table.Row

	for _, device := range devices {
		row := table.Row{
			device.Name,
			device.ID,
		}

		r = append(r, row)
	}

	return table.New(
		table.WithColumns(c),
		table.WithRows(r),
		table.WithFocused(true),
		table.WithHeight(len(devices)+1),
		table.WithStyles(table.Styles{Header: lipgloss.NewStyle().Padding(1, 0),
			Selected: lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("212"))}),
	)
}
