package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// func Sizeof(x ArbitraryType) uintptr
// func Offsetof(x ArbitraryType) uintptr
// func Alignof(x ArbitraryType) uintptr

func StringToByte(key *string) []byte {
	strPtr := (*reflect.SliceHeader)(unsafe.Pointer(key))
	strPtr.Cap = strPtr.Len
	b := *(*[]byte)(unsafe.Pointer(strPtr))
	return b
}

// uintptr <-> unsafe.Pointer <-> *type (type 表示go定义的类型)
// go 指针不能运算，但是uinitptr可以
func main() {

	changeSliceCapByUnsafePointer()

	getMapLenByUnsafePoint()

	getStructInfoUseUS()

	// decryptContent := "/AvYEjm4g6xJ3LVrk2/Adk"
	// iv := decryptContent[0:16]
	// key := decryptContent[2:18]
	// fmt.Println(&iv)
	// fmt.Println(&key)
	// ivBytes := StringToByte(&iv)
	// keyBytes := StringToByte(&key)
	// fmt.Println(string(ivBytes))
	// fmt.Println(string(keyBytes))
}

// 修改 slice cap
func changeSliceCapByUnsafePointer() {
	s := make([]int, 2, 3)
	fmt.Printf("before slice cap is %d \n", cap(s))
	// 转为 unsafe.Pointer
	ps := unsafe.Pointer(&s)
	// 转为 uintptr 进行偏移操作后再转为  unsafe.Pointer
	offsetPs := unsafe.Pointer(uintptr(ps) + uintptr(16))
	// 转为 int 赋值
	*((*int)(offsetPs)) = 6
	fmt.Printf("after slice cap is %d \n", cap(s))
}

// 获取 map len
func getMapLenByUnsafePoint() {
	m := make(map[int]int)
	m[1] = 1

	// map 首地址就是count 无需 uintptr 偏移
	len := **(**int)(unsafe.Pointer(&m))
	fmt.Printf("map len is %d \n", len)

	m[2] = 2
	len = **(**int)(unsafe.Pointer(&m))
	fmt.Printf("map len is %d \n", len)
}

// struct 内存地址是连续，内存对齐的
func getStructInfoUseUS() {
	var st = struct {
		name string
		sex  bool
		age  int
	}{"lizhi", true, 26}
	stUP := unsafe.Pointer(&st)
	fmt.Printf("sex offset = %d, age offset = %d \n", unsafe.Offsetof(st.sex), unsafe.Offsetof(st.age))
	// 获取age的值
	fmt.Printf("st.age is %d \n", *(*int)(unsafe.Pointer(uintptr(stUP) + unsafe.Offsetof(st.age))))
}

// 底层 string 转 slice
// type StringHeader struct {
// 	Data uintptr
// 	Len  int
// }

// type SliceHeader struct {
// 	Data uintptr
// 	Len  int
// 	Cap  int
// }

// func string2bytes(s string) []byte {
// 	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
// 	bh := reflect.SliceHeader{
// 	Data: stringHeader.Data,
// 	Len: stringHeader.Len,
// 	Cap: stringHeader.Len,
// 	}
// 	return *(*[]byte)(unsafe.Pointer(&bh))
// }
