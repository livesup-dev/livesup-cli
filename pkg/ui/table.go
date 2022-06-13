package ui

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

func DrawTable(modelTable ModelTable) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "First Name", "Last Name", "Email"})
	for _, row := range modelTable.GetRows() {
		t.AppendRow(table.Row{row})
		t.AppendSeparator()
	}
	t.Render()
}
