package common

import (
	"fmt"
)

func PrintStr(str any) {
	fmt.Printf("str's type is %T, str = %q \n", str, str)
}

func PrintOther(other any) {
	fmt.Printf("str's type is %T, str = %v \n", other, other)
}
