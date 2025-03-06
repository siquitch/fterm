package flows

import (
	"flutterterm/pkg/model"
	"flutterterm/pkg/ui"
	"flutterterm/pkg/utils"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type TableManager map[RunFlowStage]ui.TableModel

type RunFlowModel struct {
	devices  []model.Device
	config   model.Config
	showHelp bool

	stage RunFlowStage
	state FlowState

	runConfig    model.RunConfig
	tableManager TableManager

	spinner ui.SpinnerModel
}

type RunFlowStage int

const (
	selectDevice RunFlowStage = iota
	selectConfig
	_length
)

// Entry point to this flow
func RunFlow(config model.Config) (model.RunConfig, error) {
	p := tea.NewProgram(InitialRunModel(config))

	m, err := p.Run()

	if err != nil {
		utils.PrintError(fmt.Sprintf("Error %s", err.Error()))
	}

	rm, _ := m.(RunFlowModel)

	rc := model.RunConfig{
		SelectedConfig: rm.runConfig.SelectedConfig,
		SelectedDevice: rm.runConfig.SelectedDevice,
	}

	if err != nil {
		utils.PrintError(fmt.Sprintf("Error %s", err.Error()))
	}

	return rc, err
}

func InitialRunModel(config model.Config) RunFlowModel {
	m := RunFlowModel{
		config:       config,
		stage:        selectDevice,
		state:        getting,
		spinner:      ui.GetSpinner(),
		tableManager: make(TableManager),
	}
	return m
}

func (m RunFlowModel) Init() Cmd {
	return tea.Batch(m.spinner.Tick, getDevices())
}

func (m RunFlowModel) Update(msg Msg) (Model, Cmd) {
	switch msg := msg.(type) {

	case KeyMsg:
		if m.state == view {

			switch msg.String() {

			case "?":
				m.showHelp = !m.showHelp
				return m, nil
			case "ctrl+c", "q":
				return m, Quit
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
				switch m.stage {
				case selectConfig:
					n := m.config.Configs[m.tableManager[selectConfig].Cursor()].Name
					m.config.ToggleFavoriteConfig(n)
				}
				return m, nil
			case "enter":
				m, cmd := m.doNextThing()
				return m, cmd
			}
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
func (m *RunFlowModel) back() {
	if m.stage == selectDevice {
		return
	}
	m.stage = selectDevice
}

// Available after advancing at least once
func (m *RunFlowModel) forward() {
	if m.runConfig.SelectedDevice.ID == "" {
		return
	}
	m.stage = selectConfig
}

// Go to the next part of the process
func (m RunFlowModel) doNextThing() (RunFlowModel, Cmd) {
	var cmd Cmd
	switch m.stage {
	case selectDevice:
		m.runConfig.SelectedDevice = m.devices[m.tableManager[selectDevice].Cursor()]
		m.stage = selectConfig
		if m.tableManager[selectConfig] == nil {
			m.tableManager[selectConfig] = ui.GetConfigTable(m.config.Configs)
		}
		cmd = nil
	case selectConfig:
		m.runConfig.SelectedConfig = m.config.Configs[m.tableManager[selectConfig].Cursor()]
		cmd = tea.Quit
	}
	return m, cmd
}

func (m RunFlowModel) View() string {
	var s string = ""
	switch m.state {
	case view:
		s += fmt.Sprintf("Selected Device: %s\n", m.runConfig.SelectedDevice.Name)
		s += fmt.Sprintf("Selected Config: %s\n", m.runConfig.SelectedConfig.Name)
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
		s := fmt.Sprintf("%sGetting devices %s", spinner, spinner)
		return s
	default:
		return "Unknown state"
	}
}

func (m RunFlowModel) CurrentTable() ui.TableModel {
	return m.tableManager[m.stage]
}

func getDevices() Cmd {
	return func() Msg {
		cmd := model.FlutterDevices()
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
