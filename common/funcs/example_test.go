package funcs

import (
	"fmt"
)

func ExampleInArray() {
	var str = "4"
	var s1 = []string{"0", "1", "2"}
	i := InArray(str, s1)
	fmt.Printf("%T %v \n", i, i)
	// Output: bool false
}

// 和 php in_array 不同
func ExampleInArray_emptyString() {
	var str = ""
	var s1 = []string{"0", "1", "2"}
	i := InArray(str, s1)
	fmt.Printf("%T %v \n", i, i)
	// Output: bool false
}
