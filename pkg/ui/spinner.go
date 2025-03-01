package ui

import "github.com/charmbracelet/bubbles/spinner"

type SpinnerModel = spinner.Model
type TickMsg = spinner.TickMsg

// Default loading spinner
func GetSpinner() SpinnerModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	return s
}
