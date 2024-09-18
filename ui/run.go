package ui

import (
	"flutterterm/utils"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)


type RunModel struct {
	devices         []utils.Device
	configs         []utils.FlutterConfig
	cursor          int
	cursorLen       int
	runInfo         map[devicestage]bool
	Selected_device utils.Device
	Selected_config utils.FlutterConfig
}

type devicestage int

const (
	device devicestage = iota
	config
)

func InitialRunModel(devices []utils.Device, configs []utils.FlutterConfig) RunModel {
	return RunModel{
		devices:   devices,
		configs:   configs,
		cursor:    0,
		cursorLen: len(devices),
		runInfo:   make(map[devicestage]bool),
	}
}

func (m RunModel) Init() tea.Cmd {
	return nil
}

func (m RunModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			if m.cursor < m.cursorLen-1 {
				m.cursor++
			}

		case "enter":
			m, cmd := m.doNextThing()
			return m, cmd
		}
	}

	return m, nil
}

func (m RunModel) doNextThing() (RunModel, tea.Cmd) {
	var cmd tea.Cmd
	if _, exists := m.runInfo[device]; !exists {
		m.runInfo[device] = true
		m.Selected_device = m.devices[m.cursor]
		m.cursorLen = len(m.configs)
		cmd = nil
	} else if _, exists := m.runInfo[config]; !exists {
		m.runInfo[config] = true
		m.Selected_config = m.configs[m.cursor]
		cmd = tea.Quit
	}
	m.cursor = 0
	return m, cmd
}

// Whether the model has enough information to run
func (m RunModel) IsComplete() bool {
	return m.Selected_config.Name != "" && m.Selected_device.ID != ""
}

func (m RunModel) View() string {
	var s string
	if _, exists := m.runInfo[device]; !exists {
		// The header
		s = "Select a device\n\n"

		// Iterate over our choices
		for i, choice := range m.devices {

			// Is the cursor pointing at this choice?
			cursor := " " // no cursor
			if m.cursor == i {
				cursor = ">" // cursor!
			}

			// Render the row
			s += fmt.Sprintf("%s %s\n", cursor, choice.Name)
		}
	} else if _, exists := m.runInfo[config]; !exists {
		s = "Select a config\n\n"
		for i, config := range m.configs {
			cursor := " " // no cursor
			if m.cursor == i {
				cursor = ">" // cursor!
			}

			// Render the row
			s += fmt.Sprintf("%s %s\n", cursor, config.Name)
		}
	} else {
		s += fmt.Sprintf("Selected device: %s\n", m.Selected_device.Name)
		s += m.Selected_config.ToString()
	}

	s += "\nPress q to quit.\n"
	// Send the UI for rendering
	return s
}
