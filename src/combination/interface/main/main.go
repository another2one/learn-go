package main

import (
	"fmt"
)

// 接口是对继承的一种补充，为了规范化程序，降低耦合性

type Usb interface {
	Read()
	Write()
}

type Screen interface {
	Show()
}

type Computer struct{}

func (c *Computer) Working(usb Usb) {
	usb.Read()
	usb.Write()
}

// camera

type Camera struct{}

func (c Camera) Read() {
	fmt.Println("Camera Reading ... ")
}

func (c Camera) Write() {
	fmt.Println("Camera writting ... ")
}

func (c Camera) Show() {
	fmt.Println("Camera showing ... ")
}

// phone

type Phone struct{}

func (c *Phone) Read() {
	fmt.Println("Phone Reading ... ")
}

func (c Phone) Write() {
	fmt.Println("Phone writting ... ")
}

func main() {
	computer := Computer{}
	camera := Camera{}
	phone := Phone{}
	computer.Working(camera)
	computer.Working(&phone)

	// 细节
	// 1. interface 不能实例化，但是可以指向一个实现该接口的 “自定义类型” 变量（实例）
	var u1 Usb
	fmt.Printf("u1 = %+v, u1 = %T, u2 address is %p, camera address is %p \n", u1, u1, &u1, &camera)
	u1 = camera
	u1.Write()
	u1 = &phone // 注意，方法接收指针数据时，必须传递指针
	u1.Read()
	fmt.Printf("u1 = %+v, u1 = %T, u2 address is %p, camera address is %p \n", u1, u1, &u1, &camera)
	// 2. 一个自定义类型可以实现多个接口
	var s1 Screen
	s1 = camera

	fmt.Println(s1)

	// 3. 必须同时实现继承的其他接口才能实现该接口
	//type testInter interface {
	//	Screen
	//	Read()
	//}
	//var s2 testInter
	//s2 = camera

	s2 := interface {
		Screen
		Read()
	}(camera)

	// s2 = phone
	fmt.Println(s2)

	// 4. 不能包含相同的接口名
	// type testInter01 interface{
	// 	testInter
	// 	Usb
	// }
}
