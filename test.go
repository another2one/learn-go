package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	a := "66"
	err := errors.New("666")
	fmt.Println(err)
	var s int
	if s, err = strconv.Atoi(a); err == nil {
		fmt.Println(s, err)
	}
	fmt.Println(err)
}
