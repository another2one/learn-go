package utils

import "fmt"

func PrintSlice(s []int, tag string) {
	fmt.Printf("%s: value = %v, len = %d,  cap = %d, address = %p \n", tag, s, len(s), cap(s), s)
}
