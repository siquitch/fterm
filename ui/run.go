package ui

import (
	"flutterterm/utils"
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type RunModel struct {
	devices         []utils.Device
	configs         []utils.FlutterConfig
	cursor          utils.Cursor
	runInfo         map[devicestage]bool
	Selected_device utils.Device
	Selected_config utils.FlutterConfig
	state           state
	spinner         spinner.Model
}

type devicestage int

const (
	device devicestage = iota
	config
)

func InitialRunModel(configs []utils.FlutterConfig) RunModel {
	return RunModel{
		configs: configs,
		runInfo: make(map[devicestage]bool),
		state:   getting,
		spinner: getSpinner(),
	}
}

func (m RunModel) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, getDevices())
}

func (m RunModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			m, cmd := m.doNextThing()
			return m, cmd
		}
	case devicesComplete:
		m.devices = msg
		m.cursor = utils.NewCursor(0, len(m.devices))
		m.state = view
		return m, nil
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m RunModel) doNextThing() (RunModel, tea.Cmd) {
	var cmd tea.Cmd
	if _, exists := m.runInfo[device]; !exists {
		m.runInfo[device] = true
		m.Selected_device = m.devices[m.cursor.Index()]
		m.cursor = utils.NewCursor(0, len(m.configs))
		cmd = nil
	} else if _, exists := m.runInfo[config]; !exists {
		m.runInfo[config] = true
		m.Selected_config = m.configs[m.cursor.Index()]
		cmd = tea.Quit
	}
	return m, cmd
}

// Whether the model has enough information to run
func (m RunModel) IsComplete() bool {
	return m.Selected_config.Name != "" && m.Selected_device.ID != ""
}

func (m RunModel) View() string {

	switch m.state {
	case view:
		var s string
		if _, exists := m.runInfo[device]; !exists {
			s = "Select a device\n\n"

			for i, choice := range m.devices {
				cursor := " "
				if m.cursor.Index() == i {
					cursor = utils.CursorChar
				}
				s += fmt.Sprintf("%s %s\n", cursor, choice.Name)
			}
		} else if _, exists := m.runInfo[config]; !exists {
			s = "Select a config\n\n"
			for i, config := range m.configs {
				cursor := " "
				if m.cursor.Index() == i {
					cursor = utils.CursorChar
				}
				s += fmt.Sprintf("%s %s\n", cursor, config.Name)
			}
		} else {
			s += fmt.Sprintf("Selected device: %s\n", m.Selected_device.Name)
			s += m.Selected_config.ToString()
		}

		s += "\nPress q to quit.\n"
		return s
	case getting:
		spinner := m.spinner.View()
		s := fmt.Sprintf("%s Getting devices %s", spinner, spinner)
		return s
	default:
		return "Unknown state"
	}
}

func getDevices() tea.Cmd {
	return func() tea.Msg {
		cmd := utils.FlutterDevices()
		output, err := cmd.Output()

		if err != nil {
			return cmdError(err.Error())
		}

		devices, err := utils.ParseDevices(output)

		if err != nil {
			return cmdError(err.Error())
		}

		return devicesComplete(devices)
	}
}
