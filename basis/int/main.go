package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var n1 = 1
	var n2 int64 = 1

	// n1为int, n2为int64, 但是因为系统64位，所以大小相同
	// unsafe.Sizeof(n2) 查看占用字节数
	// fmt.Printf 格式化输出
	fmt.Printf("n1 type is %T, size is %d, n1 type is %T, int64 size is %d \n", n1, unsafe.Sizeof(n1), n2, unsafe.Sizeof(n2))

	// 转换为相同类型才能进行大小比较
	fmt.Printf("%v \n", int64(n1) == n2)

	// 最大值 math.MaxInt8
}
