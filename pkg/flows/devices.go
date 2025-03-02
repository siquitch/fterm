package flows

import (
	"flutterterm/pkg/model"
	"flutterterm/pkg/ui"
	"flutterterm/pkg/utils"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type DeviceFlowModel struct {
	devices        []model.Device
	table          ui.TableModel
	state          FlowState
	selectedDevice model.Device
	spinner        ui.SpinnerModel
	showHelp       bool
}

// Entry point for this flow
func DeviceFlow() (model.Device, error) {
	d := InitialDeviceFlowModel()

	p := tea.NewProgram(d)

	m, err := p.Run()

	if err != nil {
		utils.PrintError(fmt.Sprintf("Error %s", err.Error()))
		return model.Device{}, err
	}

	d, _ = m.(DeviceFlowModel)

	return d.selectedDevice, err
}

func InitialDeviceFlowModel() DeviceFlowModel {
	return DeviceFlowModel{
		state:   getting,
		spinner: ui.GetSpinner(),
	}
}

func (m DeviceFlowModel) Init() Cmd {
	return tea.Batch(getDevices(), m.spinner.Tick)
}

func (m DeviceFlowModel) Update(msg Msg) (Model, Cmd) {
	switch msg := msg.(type) {
	case KeyMsg:
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
			case "f":
			case "enter":
				m.selectedDevice = m.devices[m.table.Cursor()]
				return m, Quit
			}
		}
	case DevicesComplete:
		m.devices = msg
		m.state = view
		m.table = ui.GetDeviceTable(m.devices)

	case ui.TickMsg:
		var cmd Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m DeviceFlowModel) View() string {
	s := ""
	switch m.state {
	case view:
		s += "Select a device\n"
		s += m.table.View()
		s += "\n"
		s += quitAndHelpMessage
		if m.showHelp {
			s += "\n"
			s += controlsHelpMessage
		}
	case getting:
		spinner := m.spinner.View()
		s += fmt.Sprintf("%sGetting devices %s", spinner, spinner)
	}

	return s
}
