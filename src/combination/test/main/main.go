package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// 1
	var arr1 [10]int
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		arr1[i] = rand.Intn(100) + 1
	}
	fmt.Printf("arr1 = %v \n", arr1)

	var min int
	var max int
	var sum int
	for i := 0; i < 10; i++ {
		sum += arr1[i]
		if arr1[i] < arr1[min] {
			min = i
		}
		if arr1[i] > arr1[max]{
			max = i
		}
	}
	fmt.Printf("min = %v, max = %v, avg = %.1f \n", min, max, float64(sum)/10)

	for i := 0; i < 10; i++ {
		for j := i; j < 10; j++ {
			if arr1[i] < arr1[j] {
				arr1[i], arr1[j] = arr1[j], arr1[i]
			}
		}
	}
	fmt.Printf("arr1 = %v \n", arr1)
	var arr2 [11]int
	var insertInt int
	var insertIndex int
	fmt.Println("请输入插入的数 ...")
	fmt.Scanln(&insertInt)
	for index, _ := range arr1 {
		if arr1[index] < insertInt {
			insertIndex = index
			break
		}
	}
	for index, _ := range arr2 {
		if index < insertIndex {
			arr2[index] = arr1[index]
		} else if index == insertIndex {
			arr2[index] = insertInt
		} else {
			arr2[index] = arr1[index-1]
		}
	}
	fmt.Printf("arr2 = %v \n", arr2)

	arr3 :=  [3][4]int{{4, 5, 6, 7}, {7, 9, 5, 3}, {5, 5, 6, 9}}
	for _, value := range arr3 {
		fmt.Println(value)
	}
	for index, value := range arr3 {
		for index1, _ := range value {
			if index == 0 || index1 == 0 || index == 2 || index1 == 3 {
				arr3[index][index1] = 0
			}
		}
	}
	for _, value := range arr3 {
		fmt.Println(value)
	}
}