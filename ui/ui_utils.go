package ui

import (
	"flutterterm/utils"

	"github.com/charmbracelet/bubbles/spinner"
)

type devicesComplete []utils.Device

type runningComplete bool

type cmdError string

type state int

const (
	// Viewing stuff
	view state = iota
	// Loading stuff
	getting
	// Running command
	running
)

// Default loading spinner
func getSpinner() spinner.Model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	return s
}
