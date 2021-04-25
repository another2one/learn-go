package main

import (
	"fmt"
	"math"
)

func main() {

	hens := [...]float64{3.1, 3.2, 3.4, 6.5, 5.3, 4.6}

	fmt.Printf("%p, %p", &hens, &hens[1])

	// 遍历
	for index, value := range hens {
		fmt.Printf("index = %v, value = %v \n", index, value)
	}

	// 数组是值类型
	arr1 := [...]int{2, 3, 4}
	arr2 := &arr1
	arr2[0] = 3
	fmt.Printf("arr1 = %v, arr2 = %v \n", arr1, arr2)

	// 输出a-z的数组
	byteArr := [26]byte{'A'}
	for index, _ := range byteArr {
		if index > 0 {
			byteArr[index] = byteArr[index-1] + 1
		}
	}
	fmt.Printf("byteArr = %q \n", byteArr)

	// 数组反转
	count := len(byteArr)
	for index, _ := range byteArr {
		if index < int(math.Ceil(float64(count)/2)) {
			byteArr[index], byteArr[count-index-1] = byteArr[count-index-1], byteArr[index]
		}
	}
	fmt.Printf("byteArr = %q \n", byteArr)

	// 转为slice
	sl := byteArr[:]
	fmt.Printf("sl = %q \n", sl)

	// array转为string
	// 1
	str := fmt.Sprintf("%q \n", byteArr)
	fmt.Println("str: ", str)

	var a1 = [...]int8{1, 2, 3, 4, 5}
	fmt.Printf("a1 address = %p\n", &a1)

	// 数组反转`
	fmt.Printf("a1[2:] type is %T address = %p\n", a1[2:], a1[2:])
	a1Len := len(a1)
	for start := 0; start <= int(math.Floor(float64(a1Len)/2)); {
		a1[start], a1[a1Len-start-1] = a1[a1Len-start-1], a1[start]
		start++
	}
	fmt.Printf("%v reverse to %v \n", [...]int8{1, 2, 3, 4, 5}, a1)
}
