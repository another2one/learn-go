package main

import (
	"fmt"
)

func main() {
	var student = [6][4]int{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}
	for index1, value1 := range student {
		for index2, _ := range value1 {
			fmt.Printf("student[%d][%d] address is %p \n", index1, index2, &student[index1][index2])
		}
	}
}