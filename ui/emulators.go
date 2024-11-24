package ui

import (
	"flutterterm/utils"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type EmulatorModel struct {
	devices          []utils.Device
	cursor           utils.Cursor
	SelectedEmulator utils.Device
}

func InitialEmulatorModel(devices []utils.Device) EmulatorModel {
	return EmulatorModel{
		devices: devices,
		cursor:  utils.NewCursor(0, len(devices)),
	}
}

func (m EmulatorModel) Init() tea.Cmd {
	return nil
}

func (m EmulatorModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			m.cursor.Previous()
		case "down", "j":
			m.cursor.Next()
		case "enter":
			m.SelectedEmulator = m.devices[m.cursor.Index()]
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m EmulatorModel) View() string {
	var s string

	s = "Select an emulator\n\n"

	for i, device := range m.devices {
		cursor := " "
		if m.cursor.Index() == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, device.Name)
	}

	s += "\nPress q to quit\n"

	return s
}

func (m EmulatorModel) IsComplete() bool {
	return m.SelectedEmulator.ID != "" && m.SelectedEmulator.Name != ""
}
