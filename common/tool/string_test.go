package tool

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"testing"
)

func TestGetCell(t *testing.T) {
	type args struct {
		rowIndex    int
		columnIndex int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"test1", args{0, 0}, "A1"},
		{"test1", args{25, 0}, "Z1"},
		{"test3", args{26, 0}, "AA1"},
		{"test4", args{52, 0}, "BA1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCell(tt.args.rowIndex, tt.args.columnIndex); got != tt.want {
				t.Errorf("GetCell() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateHexColor(t *testing.T) {
	var style = lipgloss.NewStyle().Bold(true)
	for i := 0; i < 100; i++ {
		color := GenerateHexColor()
		style.Foreground(lipgloss.Color(color))
		fmt.Println(style.Render(color))
	}
}
