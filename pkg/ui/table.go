package ui

import (
	"fterm/pkg/model"

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

func GetDeviceTable(devices []model.Device) TableModel {
	c := []TableColumn{
		{Title: "Name", Width: 20},
		{Title: "ID", Width: 30},
	}

	var r []TableRow

	for _, device := range devices {
		row := TableRow{
			device.Name,
			device.ID,
		}

		r = append(r, row)
	}

	t := GetTable(c, r)

	return t
}

func GetConfigTable(configs []model.FlutterConfig) TableModel {
	c := []TableColumn{
		{Title: "", Width: 2},
		{Title: "Config", Width: 20},
		{Title: "Description", Width: 30},
	}

	var r []TableRow

	for _, config := range configs {
        s := ""

		if config.Favorite {
            s = starIcon
		}
		row := TableRow{s, config.Name, config.Name}

		r = append(r, row)
	}

	t := GetTable(c, r)

	return t
}
