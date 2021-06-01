package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func StringToByte(key *string) []byte {
	// error: 不能强转类型，会造成覆写
	strPtr := (*reflect.SliceHeader)(unsafe.Pointer(key))
	strPtr.Cap = strPtr.Len
	b := *(*[]byte)(unsafe.Pointer(strPtr))
	return b
}

// 重写方法 底层目前也是这么实现的
func NewStringToByte(key *string) []byte {
	strPtr := (*reflect.StringHeader)(unsafe.Pointer(key))
	nb := reflect.SliceHeader{
		Data: strPtr.Data,
		Len:  strPtr.Len,
		Cap:  strPtr.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&nb))
}

// https://mp.weixin.qq.com/s/tXAP8_U63QLNj1h0ZMvXPw
// 变量在栈中顺序排，直接操作地址可能发生覆写
func main() {
	decryptContent := "/AvYEjm4g6xJ3LVrk2/Adk"

	// 变量在栈上 每个占 8 字节 type stringHeader struct{Data uintptr, Len int}
	// len
	// Data <--- key
	// len
	// Data  <--- iv
	iv := decryptContent[0:16]
	key := decryptContent[1:18]

	fmt.Println(&iv)
	fmt.Println(&key)

	// ??? 先转key就正常 先传 iv 导致 key 的 *str 被cap覆盖
	// type sliceHeader struct{Data uintptr, Len int, Cap int} 多的cap占了栈上key.Data的位置，讲key.Data变为16，原本为指向字符串的地址
	// 如果先处理key就不会发生覆写
	// ivBytes := StringToByte(&iv)
	// keyBytes := StringToByte(&key)

	// 重写转换方法
	ivBytes := NewStringToByte(&iv)
	keyBytes := NewStringToByte(&key)

	fmt.Println(string(ivBytes))
	fmt.Println(string(keyBytes))
}
