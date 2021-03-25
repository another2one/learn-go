package main

import (
	"fmt"
	"unsafe"
)

func main() {

	// var c1 byte = 'c'
	// var c2 rune = '你' // utf-8 占用3字节，无法用byte表示
	c1 := 'c'
	c2 := '你'
	c3 := []byte("cc")
	c4 := []byte("c1你")

	fmt.Printf("c1 type is %T, size is %d, c2 type is %T, size is %d, c3 type is %T, size is %d, c4 type is %T, size is %d \n",
		c1, unsafe.Sizeof(c1), c2, unsafe.Sizeof(c2), c3, len(c3), c4, len(c4))
	fmt.Println("c1 =", c1, "\t c2 =", c2)
	fmt.Printf("c1 = %c \t c2 = %c \n", c1, c2)
}
