package main

import (
	"fmt"

	"github.com/duke-git/lancet/v2/strutil"
)

func main() {
	s := "hello"
	rs := strutil.ReverseStr(s)
	fmt.Println(rs) //olleh
}
