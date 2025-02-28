package utils

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

type TableModel = *table.Model
type TableRow = table.Row
type TableColumn = table.Column

func GetTable(c []TableColumn, r []TableRow) TableModel {
	t := table.New(
		table.WithColumns(c),
		table.WithRows(r),
		table.WithHeight(len(r)+1),
		table.WithStyles(table.Styles{
			Header:   lipgloss.NewStyle().Padding(1, 0),
			Cell:     lipgloss.Style{},
			Selected: lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("212")),
		}),
	)
	return &t
}
