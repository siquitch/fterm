package ui

import (
	// "github.com/charmbracelet/bubbles/cursor"
	"fmt"

	"github.com/charmbracelet/bubbles/textarea"
	// "github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type CreateEmulatorModel struct {
	textarea textarea.Model
	text     string
	state    modelState
}

type modelState int

const (
	inProgress modelState = iota
	complete
	abandoned
)

func InitialCreateEmulatorModel() CreateEmulatorModel {
	ta := textarea.New()
	ta.Placeholder = "Enter a name for your emulator"
	ta.Focus()
	ta.Prompt = "| "
	ta.CharLimit = 20
	ta.SetWidth(25)
	ta.SetHeight(1)
	ta.ShowLineNumbers = false
	return CreateEmulatorModel{
		textarea: ta,
		state:    inProgress,
	}
}

func (m CreateEmulatorModel) Text() string {
	return m.text
}

func (m CreateEmulatorModel) IsComplete() bool {
	return m.state == complete
}

func (m CreateEmulatorModel) Init() tea.Cmd {
	return textarea.Blink
}

func (m CreateEmulatorModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.textarea.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {

		case "esc", "ctrl+c":
			// Quit
			m.state = abandoned
			return m, tea.Quit
		case "enter":
			v := m.textarea.Value()
			// Don't allow empty input
			if v == "" {
				return m, nil
			}
			m.text = m.textarea.Value()
			m.state = complete
			return m, tea.Quit
		default:
			var cmd tea.Cmd
			m.textarea, cmd = m.textarea.Update(msg)
			return m, cmd
		}
	default:
		return m, nil
	}
}

func (m CreateEmulatorModel) View() string {
	switch m.state {
	case inProgress:
		return fmt.Sprintf("> %s\n", m.textarea.View())
	case abandoned:
		return ""
	default:
		return fmt.Sprintf("Creating %s", m.text)
	}
}
