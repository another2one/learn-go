package main

import (
	"fmt"
)

// https://mp.weixin.qq.com/s?__biz=Mzg5NjIwNzIxNQ==&mid=2247484072&idx=2&sn=0363c7102943888e0f390f3f5a9ae662&chksm=c005d2a8f7725bbe068f418f72bd9f8ecc3aec3daf11ba7e7ad80d202e6ad90ceec459c5c7e8&scene=21#wechat_redirect
// 非空接口底层
// type iface struct {
// 	tab  *itab // interface 信息
// 	data unsafe.Pointer // 动态数据
// }

// type itab struct {
// 	inter *interfacetype // 底层优化会做缓存 interface相关信息
// 	_type *_type // 接口的动态类型，也就是被赋给接口类型的那个变量的类型元数据
// 	hash uint32 // _type的hash值 用于快速判断interface类型是否相等
// 	_ [4]byte
// 	fun [n]uintptr // 指向 _type 的实现方法地址 如果itab._type对应的类型没有实现这个接口，则itab.fun[0]=0
// }

// type interfacetype struct {
// 	typ _type
// 	pkgpath nametype // 位置
// 	mhdr []imethod // 接口方法slice
// }
//
// 空接口底层
// type eface struct {
// 	 _type *_type // data的数据类型
//   data unsafe.Pointer
// }
type w interface {
	say()
}

type w1 struct{}

func (w1 *w1) say() {
	fmt.Println("say")
}

func (w1 *w1) said() {
	fmt.Println("said")
}

func main() {
	var e1 w // 此时 iface.data = nil
	s1 := w1{}
	e1 = &s1 // iface.data 指向 &s1
	e1.say()
	// e1.said() // error
}
