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
	cursor          utils.Navigator
	stage           devicestage
	Selected_device utils.Device
	Selected_config utils.FlutterConfig
	state           state
	spinner         spinner.Model
}

type devicestage int

const (
	device devicestage = iota
	config
	_length
)

func InitialRunModel(configs []utils.FlutterConfig) RunModel {
	return RunModel{
		configs: configs,
		stage:   device,
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

		case "?":
			m.cursor.ToggleHelp()

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			m.cursor.Previous()

		case "down", "j":
			m.cursor.Next()

		case "left", "h":
			m = m.back()

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

// Go back in the process
func (m RunModel) back() RunModel {
	if m.stage == device {
		return m
	}
	m.stage = device
	return m
}

// Go to the next part of the process
func (m RunModel) doNextThing() (RunModel, tea.Cmd) {
	var cmd tea.Cmd
	switch m.stage {
	case device:
		m.Selected_device = m.devices[m.cursor.Index()]
		m.cursor.Reset(len(m.configs))
		m.stage = config
		cmd = nil
	case config:
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
	var s string = ""
	if m.cursor.ShouldShowHelp() {
		s += controlsHelpMessage
	}
	switch m.state {
	case view:
		switch m.stage {
		case device:
			s += "Select a device\n\n"

			for i, device := range m.devices {
				cursor := " "
				if m.cursor.Index() == i {
					cursor = utils.CursorChar
				}
				s += fmt.Sprintf("%s %s - %s\n", cursor, device.Name, device.ID)
			}
		case config:
			s += "Select a config\n\n"
			for i, config := range m.configs {
				cursor := " "
				if m.cursor.Index() == i {
					cursor = utils.CursorChar
				}
				s += fmt.Sprintf("%s %s\n", cursor, config.Name)
			}
		}

		s += fmt.Sprintf("\n%d/%d", m.stage+1, _length)
		s += quitAndHelpMessage
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
