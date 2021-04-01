package utils

import "fmt"

func PrintSlice(s []int, tag string) {
	fmt.Printf("%s: value = %v, len = %d,  容量 = %d \n", tag, s, len(s), cap(s))
}
