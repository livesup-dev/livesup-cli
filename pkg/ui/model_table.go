package ui

type ModelTable interface {
	GetHeaders() []string
	GetRows() [][]string
}
