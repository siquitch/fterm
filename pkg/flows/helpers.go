package flows

import (
	"fterm/pkg/model"

	tea "github.com/charmbracelet/bubbletea"
)

type DevicesComplete []model.Device

type RunningComplete bool

type CmdError string

type FlowState int

const (
	// Viewing stuff
	view FlowState = iota
	// Loading stuff
	getting
	// Running command
	running
)

type Model = tea.Model
type Cmd = tea.Cmd
type Msg = tea.Msg
type KeyMsg = tea.KeyMsg

func Quit() Msg {
	return tea.QuitMsg{}
}

const quitAndHelpMessage = "\nPress q to quit, or ? for help\n"
const controlsHelpMessage = "Controls\nj, down: go down\nk, up: go up\nh, left: go back (if applicable)\nenter: submit\n\n"
