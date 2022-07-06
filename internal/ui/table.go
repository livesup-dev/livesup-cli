package ui

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

func DrawTable(modelTable ModelTable) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	t.AppendHeader(buildRowFromSlice(modelTable.GetHeaders()))
	t.SetStyle(table.StyleLight)
	// t.SetColumnConfigs([]table.ColumnConfig{
	// 	{Name: "#", WidthMax: 20, WidthMaxEnforcer: text.WrapSoft},
	// })

	for _, row := range modelTable.GetRows() {
		t.AppendRow(buildRowFromSlice(row))
	}
	t.Render()
}

func buildRowFromSlice(slice []string) table.Row {
	// Im not so sure yet why we need this
	// https://github.com/jedib0t/go-pretty/issues/201
	var row table.Row
	for _, element := range slice {
		row = append(row, element)
	}

	return row
}
