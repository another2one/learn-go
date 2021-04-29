package main

import (
	"fmt"
	"runtime"
)

// go 切片引用底层数组造成的内存泄漏
func main() {

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("start: %v \n", m.Sys)

	type node [100]byte
	var nodeSlice []*node
	for i := 0; i < 10; i++ {
		nodeSlice = append(nodeSlice, &node{})
	}

	nodeSlice = nodeSlice[:5]
	// 此时nodeslice底层数组 5-9 不会被释放
	// 需要先清空
	// for i := 5; i < len(nodeSlice); i++ {
	// 	nodeSlice[i] = nil
	// }
	// nodeSlice = nodeSlice[:5]

	runtime.ReadMemStats(&m)
	fmt.Printf("after: %v \n", m.Sys)

	fmt.Printf("%v \n", nodeSlice)
}
