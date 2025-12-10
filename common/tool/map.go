package tool

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func MapPrint[T comparable, V any](tag string, m map[T]V) {
	table := table.New().Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
		StyleFunc(func(row, col int) lipgloss.Style {
			return lipgloss.NewStyle().Padding(0, 2).Foreground(lipgloss.Color("99"))
		})
	for k, v := range m {
		table.Row(fmt.Sprintf("%v", k), fmt.Sprintf("%+v", v))
	}

	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		Align(lipgloss.Center).
		Width(24)

	fmt.Println(style.Render(tag))
	fmt.Print(table)
}
