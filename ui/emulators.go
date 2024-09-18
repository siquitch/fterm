package ui

import (
	"flutterterm/utils"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type EmulatorModel struct {
	devices          []utils.Device
	cursor           int
	SelectedEmulator utils.Device
}

func InitialEmulatorModel(devices []utils.Device) EmulatorModel {
	return EmulatorModel{
		devices: devices,
		cursor:  0,
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
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.devices)-1 {
				m.cursor++
			}
		case "enter":
			m.SelectedEmulator = m.devices[m.cursor]
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
		if m.cursor == i {
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
