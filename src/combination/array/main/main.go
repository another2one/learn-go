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
	arr2[0]= 3
	fmt.Printf("arr1 = %v, arr2 = %v \n", arr1, arr2)

	byteArr := [26]byte{'A'}

	for index, _ := range byteArr {
		if index > 0 {
			byteArr[index] = byteArr[index-1] + 1
		}
	}
	fmt.Printf("byteArr = %q", byteArr)

	count := len(byteArr)
	for index, _ := range byteArr {
		if index < int(math.Ceil(float64(count/2))) {
			byteArr[index], byteArr[count-index-1] = byteArr[count-index-1], byteArr[index]
		}
	}
	fmt.Printf("byteArr = %q", byteArr)
}