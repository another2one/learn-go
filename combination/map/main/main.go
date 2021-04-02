package main

import (
	"fmt"
)

func main() {

	// **********************  几种定义使用方法 ************************

	// 第一种 make分配空间后才能使用
	var m1 map[string]string
	// m1["我是谁"] = "李志" //报错 panic: assignment to entry in nil map
	m1 = make(map[string]string, 2)
	m1["我是谁"] = "李志"
	fmt.Printf("%v, %d \n", m1, len(m1))

	// 第二种
	m2 := make(map[string]string, 2)
	m2["我是谁"] = "李志"
	fmt.Printf("%v, %d \n", m2, len(m2))

	// 第三种
	m3 := map[string]string{
		"我是谁":   "李志",
		"lizhi": "ll",
	}
	fmt.Printf("%v, %d \n", m3, len(m3))

	// **********************  增删改查 ************************

	m3["sex"] = "男"
	fmt.Printf("%v, %d \n", m3, len(m3))
	delete(m3, "sex")
	// m3 = make(map[string]string)  // 全删除 gc回收
	fmt.Printf("%v, %d \n", m3, len(m3))
	if _, ok := m3["666"]; !ok {
		fmt.Println("not found")
	}

	// **********************  mapSlice ************************

	mapSlice := make([]map[string]string, 2)
	fmt.Printf("%v, %T \n", mapSlice, mapSlice)

	testMap := map[string]map[string]string{
		"lizhi": map[string]string{
			"nickName": "ss",
			"passWord": "666",
		},
	}
	if testMap["lizhi"] != nil {
		testMap["lizhi"]["passWord"] = "8888888"
	}
	fmt.Printf("%v, %T \n", testMap, testMap)

	// **********************  map 为引用类型 ************************

	m11 := map[int]int{1: 1, 2: 1}
	change(m11)
	fmt.Printf("%+v, %T \n", m11, m11)
}

func change(ss map[int]int) {
	ss[2] = 4
}
