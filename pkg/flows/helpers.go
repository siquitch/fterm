package flows

import (
	"flutterterm/pkg/model"
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

// Unicode for the star icon
const starIcon = "\u2B50"


const quitAndHelpMessage = "\nPress q to quit, or ? for help\n"
const controlsHelpMessage = "Controls\nj, down: go down\nk, up: go up\nh, left: go back (if applicable)\nenter: submit\n\n"
