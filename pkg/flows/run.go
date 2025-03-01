package flows

import (
	"flutterterm/pkg/command"
	"flutterterm/pkg/model"
	"flutterterm/pkg/ui"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type SelectionManager struct {
	SelectedDevice model.Device
	SelectedConfig model.FlutterConfig
}

type TableManager map[deviceStage]ui.TableModel

type RunModel struct {
	devices  []model.Device
	config   model.Config
	showHelp bool

	stage deviceStage
	state FlowState

	selectionManager SelectionManager
	tableManager     TableManager

	spinner ui.SpinnerModel
}

type deviceStage int

const (
	selectDevice deviceStage = iota
	selectConfig
	_length
)

type Model = tea.Model
type Cmd = tea.Cmd
type Msg = tea.Msg
type KeyMsg = tea.KeyMsg

func InitialRunModel(config model.Config) RunModel {
	m := RunModel{
		config:       config,
		stage:        selectDevice,
		state:        getting,
		spinner:      ui.GetSpinner(),
		tableManager: make(TableManager),
	}
	return m
}

func (m RunModel) Init() Cmd {
	return tea.Batch(m.spinner.Tick, getDevices())
}

func (m RunModel) Update(msg Msg) (Model, Cmd) {
	switch msg := msg.(type) {

	case KeyMsg:

		switch msg.String() {

		case "?":
			m.showHelp = !m.showHelp
			return m, nil
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.tableManager[m.stage].Cursor() == 0 {
				m.tableManager[m.stage].GotoBottom()
			} else {
				m.tableManager[m.stage].MoveUp(1)
			}
		case "down", "j":
			if m.tableManager[m.stage].Cursor()+1 >= len(m.tableManager[m.stage].Rows()) {
				m.tableManager[m.stage].GotoTop()
			} else {
				m.tableManager[m.stage].MoveDown(1)
			}
		case "left", "h":
			m.back()
			return m, nil
		case "right", "l":
			m.forward()
			return m, nil
		case "f":
		case "enter":
			m, cmd := m.doNextThing()
			return m, cmd
		}
		return m, nil

	case DevicesComplete:
		m.devices = msg
		m.state = view
		m.tableManager[selectDevice] = ui.GetDeviceTable(m.devices)
		return m, nil
	case ui.TickMsg:
		var cmd Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	return m, nil
}

// Go back in the process
func (m *RunModel) back() {
	if m.stage == selectDevice {
		return
	}
	m.stage = selectDevice
}

// Available after advancing at least once
func (m *RunModel) forward() {
	if m.selectionManager.SelectedDevice.ID == "" {
		return
	}
	m.stage = selectConfig
}

// Go to the next part of the process
func (m RunModel) doNextThing() (RunModel, Cmd) {
	var cmd Cmd
	switch m.stage {
	case selectDevice:
		m.selectionManager.SelectedDevice = m.devices[m.tableManager[selectDevice].Cursor()]
		m.stage = selectConfig
		if m.tableManager[selectConfig] == nil {
			m.tableManager[selectConfig] = ui.GetConfigTable(m.config.Configs)
		}
		cmd = nil
	case selectConfig:
		m.selectionManager.SelectedConfig = m.config.Configs[m.tableManager[selectConfig].Cursor()]
		cmd = tea.Quit
	}
	return m, cmd
}

// Whether the model has enough information to run
func (m RunModel) IsComplete() bool {
	return m.selectionManager.SelectedConfig.Name != "" && m.selectionManager.SelectedDevice.ID != ""
}

func (m RunModel) SelectedDevice() model.Device {
	return m.selectionManager.SelectedDevice
}

func (m RunModel) SelectedConfig() model.FlutterConfig {
	return m.selectionManager.SelectedConfig
}

func (m RunModel) View() string {
	var s string = ""
	switch m.state {
	case view:
		s += fmt.Sprintf("Selected Device: %s\n", m.selectionManager.SelectedDevice.Name)
		s += fmt.Sprintf("Selected Config: %s\n", m.selectionManager.SelectedConfig.Name)
		s += m.tableManager[m.stage].View()
		s += "\n"
		s += fmt.Sprintf("\n%d/%d", m.stage+1, _length)
		s += quitAndHelpMessage

		if m.showHelp {
			s += "\n"
			s += controlsHelpMessage
		}
		return s
	case getting:
		spinner := m.spinner.View()
		s := fmt.Sprintf("%s Getting devices %s", spinner, spinner)
		return s
	default:
		return "Unknown state"
	}
}

func (m RunModel) CurrentTable() ui.TableModel {
	return m.tableManager[m.stage]
}

func getDevices() Cmd {
	return func() Msg {
		cmd := command.FlutterDevices()
		output, err := cmd.Output()

		if err != nil {
			return CmdError(err.Error())
		}

		devices, err := model.ParseDevices(output)

		if err != nil {
			return CmdError(err.Error())
		}

		return DevicesComplete(devices)
	}
}
