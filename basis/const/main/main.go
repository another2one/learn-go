package main

import "fmt"

type Week int64

// const (
// 	Monday Week = iota + 1
// 	Tuesday
// )

// const (
// 	_ Week = iota
// 	Monday
// 	Tuesday
// )

const (
	_       Week = iota
	Monday       = 1 << (10 * iota) // 10 << 10
	Tuesday                         // 10 << 20
)

// const (
// 	a = iota // 0
// 	b = 5
// 	c = iota // 2
// 	d = 6    // 6
// 	e        // 6 打断后没有iota再接上 后面全为6
// 	f
// )

const (
	a = iota // 0
	b = 5
	c = iota // 2
	d = 6    // 6
	e = iota // 4 接上后继续开始递增
	f        // 5
)

func main() {
	fmt.Println("Moday = ", Monday)
	fmt.Println("Tuesday = ", Tuesday)
	fmt.Println([...]int{a, b, c, d, e, f})
}
