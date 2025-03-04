package flows

import (
	"flutterterm/pkg/model"
	"flutterterm/pkg/ui"
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type runConfigInput int

const (
	name runConfigInput = iota
	desc
	mode
	flavor
	target
	define
	additional
	_totalInputs
)

var fcInputs = map[runConfigInput]string{
	name:       "Name",
	desc:       "Description",
	mode:       "Mode",
	flavor:     "Flavor",
	target:     "Target",
	define:     "Dart Define File",
	additional: "Additional args",
}

type inputMap = map[runConfigInput]*ui.TextInput

type AddFlowModel struct {
	config       model.FlutterConfig
	currentInput runConfigInput
	viewport     ui.ViewportModel
	inputs       inputMap
}

func AddFlow() (model.FlutterConfig, error) {
	p := tea.NewProgram(InitialAddFlowModel())

	m, err := p.Run()

	fm, _ := m.(AddFlowModel)

	return fm.config, err
}

func InitialAddFlowModel() AddFlowModel {
	fm := AddFlowModel{
		currentInput: name,
		inputs:       make(inputMap),
	}

	for k, v := range fcInputs {
		ti := textinput.New()
		ti.Prompt = fmt.Sprintf("%s: ", v)
		fm.inputs[k] = &ti
	}

	fm.inputs[fm.currentInput].Focus()

	return fm
}

func (m AddFlowModel) Init() Cmd {
	return textinput.Blink
}

func (m AddFlowModel) Update(msg Msg) (Model, Cmd) {
	switch msg := msg.(type) {
	case ui.WindowSizeMsg:
		m.viewport.Height = msg.Height
		m.viewport.Width = msg.Width
		return m, nil
	case KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, Quit
		case "down", tea.KeyTab.String():
			m.focusNext()
			return m, nil
		case "up", tea.KeyShiftTab.String():
			m.focusPrevious()
			return m, nil
		case "enter":
			return m, Quit
		default:
			var cmd Cmd
			*m.inputs[m.currentInput], cmd = m.inputs[m.currentInput].Update(msg)
			return m, cmd
		}
	}
	return m, nil
}

func (m AddFlowModel) currentInputModel() *ui.TextInput {
	return m.inputs[m.currentInput]
}

func (m *AddFlowModel) focusNext() {
	m.currentInputModel().Blur()
	m.currentInput = m.currentInput + 1
	if m.currentInput == _totalInputs {
		m.currentInput = 0
	}
	m.currentInputModel().Focus()
}

func (m *AddFlowModel) focusPrevious() {
	m.currentInputModel().Blur()
	m.currentInput = m.currentInput - 1
	if m.currentInput < 0 {
		m.currentInput = _totalInputs - 1
	}
    m.currentInputModel().Focus()
}

func (m AddFlowModel) View() string {
	s := ""

	for key := range len(fcInputs) {
		s += m.inputs[runConfigInput(key)].View()
		s += "\n"
	}

	return s
}
