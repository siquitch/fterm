package flows

import (
	"fmt"
	"fterm/pkg/model"
	"fterm/pkg/ui"
	"fterm/pkg/utils"

	tea "github.com/charmbracelet/bubbletea"
)

type EmulatorFlowModel struct {
	devices          []model.Device
	selectedEmulator model.Device
	state            FlowState
	spinner          ui.SpinnerModel
	coldStart        bool // Cold start
	table            ui.TableModel
	showHelp         bool
	config           model.Config
}

func EmulatorFlow(config model.Config, isCold bool) error {
	p := tea.NewProgram(InitialEmulatorModel(config, isCold))

	_, err := p.Run()

	return err
}

func InitialEmulatorModel(config model.Config, isCold bool) EmulatorFlowModel {
	return EmulatorFlowModel{
		state:   getting,
		spinner: ui.GetSpinner(),
		config:  config,
	}
}

func (m EmulatorFlowModel) Init() Cmd {
	return tea.Batch(m.spinner.Tick, getEmulators(m.config.Fvm))
}

func (m EmulatorFlowModel) Update(msg Msg) (Model, Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.state == view {
			switch msg.String() {
			case "?":
				m.showHelp = !m.showHelp
			case "ctrl+c", "q":
				return m, Quit
			case "up", "k":
				if m.table.Cursor() == 0 {
					m.table.GotoBottom()
				} else {
					m.table.MoveUp(1)
				}
			case "down", "j":
				if m.table.Cursor()+1 >= len(m.table.Rows()) {
					m.table.GotoTop()
				} else {
					m.table.MoveDown(1)
				}
			case "enter":
				m.selectedEmulator = m.devices[m.table.Cursor()]
				m.state = running
				return m, m.launchEmulator()
			}
		}
	case ui.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case DevicesComplete:
		m.state = view
		m.devices = msg
		m.table = ui.GetDeviceTable(m.devices)
		return m, nil
	case RunningComplete:
		return m, tea.Quit
	case CmdError:
		utils.PrintError(fmt.Sprintf("%s", msg))
		m.state = view
		return m, tea.Quit
	}
	return m, nil
}

func (m EmulatorFlowModel) View() string {
	var s string = ""
	switch m.state {
	case view:
		s += "Select an emulator\n"
		s += m.table.View()
		s += quitAndHelpMessage
		if m.showHelp {
			s += "\n\n"
			s += controlsHelpMessage
		}
		return s
	case getting:
		spinner := m.spinner.View()
		return fmt.Sprintf("%sGetting emulators %s", spinner, spinner)
	case running:
		spinner := m.spinner.View()
		return fmt.Sprintf("%sLaunching %s %s", spinner, m.selectedEmulator.Name, spinner)
	default:
		return "Unknown state"
	}

}

func getEmulators(fvm bool) tea.Cmd {
	return func() tea.Msg {
		cmd := model.FlutterEmulators(fvm)

		output, err := cmd.Output()

		if err != nil {
			return CmdError(err.Error())
		}

		devices, err := model.ParseEmulators(output)
		if err != nil {
			return CmdError(err.Error())
		}

		return DevicesComplete(devices)
	}
}

func (m EmulatorFlowModel) launchEmulator() Cmd {
	return func() Msg {
		cmd := m.selectedEmulator.BuildLaunchEmulatorCommand(m.config, m.coldStart)
		err := cmd.Run()
		if err != nil {
			return CmdError(err.Error())
		}
		return RunningComplete(true)
	}
}
