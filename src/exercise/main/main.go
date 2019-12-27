package main

import(
	_"fmt"
	"common"
	"strconv"
)

func main() {
	str, _ := strconv.ParseInt("10", 10, 64)
	common.PrintOther(str)
	str2, _ := strconv.ParseFloat("1.23", 64)
	common.PrintOther(str2)
	str3, _ := strconv.ParseBool("true")
	common.PrintOther(str3)
}