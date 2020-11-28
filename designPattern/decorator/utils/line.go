package utils

import "fmt"

type Line struct {
}

func (line Line) Draw() string {
	fmt.Println("直线")
	return "直线"
}
