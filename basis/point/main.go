package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	var i int = 65

	// 指针定义 取地址 &
	var p *int = &i
	// var p *int  // 空指针值为nil
	print(p)
	// 指针取值 *
	print(*p)

	*p = 66
	print(i)

	var _ int = 64

	// https://learnku.com/articles/39255
	// uintptr(unsafe.Pointer(w)) 获取了 w 的指针起始值，
	// unsafe.Offsetof(w.b) 获取 b 变量的偏移量 (因为内存对其，占用字节数为8的倍数，不足会补齐)
	// 1. 平台移植性：32 位平台上运行 64 位平台上编译的程序要求必须 8 字节对齐 2. 寻址优化
	// 两个相加就得到了 b 的地址值，将通用指针 Pointer 转换成具体指针 ((*int)(b))，
	// 通过 符号取值，然后赋值，(*int)(b) 相当于把 b 转换成 *int 了，最后对变量重新赋值成 10，这样指针运算就完成了。
	type W struct {
		e [5]bool // 0
		a int8    // 5
		b int32   // 8
		c int64   // 16
	}
	var w *W = new(W)
	//这时w的变量打印出来都是默认值0，0
	fmt.Println(w.b, w.c)

	//现在我们通过指针运算给b变量赋值为10
	fmt.Printf("struct 中 a(%T)的偏移量: %d \n", w.a, unsafe.Offsetof(w.a))
	fmt.Printf("struct 中 b(%T)的偏移量: %d \n", w.b, unsafe.Offsetof(w.b))
	fmt.Printf("struct 中 c(%T)的偏移量: %d \n", w.c, unsafe.Offsetof(w.c))
	b := unsafe.Pointer(uintptr(unsafe.Pointer(w)) + unsafe.Offsetof(w.b))
	*((*int)(b)) = 10
	//此时结果就变成了10，0
	fmt.Println(w.b, w.c)

	// string
	// type stringStruct struct {
	// 	str unsafe.Pointer // 字符串首地址，只读
	// 	len int // 长度
	// }
	s := "sss"
	sp := unsafe.Pointer(&s)
	fmt.Println(*((*string)(sp)))
	sl := unsafe.Pointer(uintptr(sp) + 8)
	fmt.Println(*((*int)(sl)))
	*((*int)(sl)) = 6                                                                // 危险操作
	fmt.Printf("s = %v, len = %d, unsafe.size = %d \n", s, len(s), unsafe.Sizeof(s)) // 666

	// slice
	// type slice struct {
	//     array unsafe.Pointer
	//     len   int
	//     cap   int
	// }

	slice1 := []int{0, 0, 1}
	slicePointer := unsafe.Pointer(&slice1)
	print(slicePointer)
	sss := (*[]int)(slicePointer)
	(*sss)[1] = 5
	print((*sss)[1])
	print(slice1)
	sr := reflect.ValueOf(slicePointer)
	if sr.CanSet() {
		fmt.Println("sr can be set")
	} else {
		fmt.Println("sr can't be set")
	}
}

func print(val interface{}) {
	fmt.Printf("val's type is %T, val = %v \n", val, val)
}
