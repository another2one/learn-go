package main

import (
	"fmt"
	"unsafe"
)

func main()  {

	var n1 = 1.1;
	var n2 float64 = 1;

	// n1,n2都为float64, 和整型不同
	fmt.Printf("n1 type is %T, size is %d, n1 type is %T, int64 size is %d \n", n1, unsafe.Sizeof(n1), n2, unsafe.Sizeof(n2));

	// 精度损失
	var n3 float32 = 123.0000901
	var n4 float64 = 123.0000901
	fmt.Println("n3 =", n3, "n4 =", n4);

	// 表示
	var n5 float32 = 123.0000901
	var n6 float64 = 9.01e-5
	fmt.Println("n5 =", n5, "n6 =", n6);
}