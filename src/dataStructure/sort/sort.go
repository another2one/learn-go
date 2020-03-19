package main

import (
	"fmt"
	"math/rand"
	"time"
)

const NUM = 1000

func getArray1() func() [NUM]int {
	var arr [NUM]int
	rand.Seed(time.Now().Unix())
	for i := 0; i < NUM; i++ {
		arr[i] = rand.Intn(NUM)
		//arr = [NUM]int{15, 29, 10, 15, 28, 50, 21, 20, 24, 28, 23}
	}
	return func() [NUM]int {
		return arr
	}
}

var (
	getArray = getArray1()
	timeArr  [2]int
	flag     = 0
)

func useTime() {
	if flag == 1 {
		timeArr[flag] = time.Now().Second()
		fmt.Println(timeArr[1] - timeArr[0])
		flag = 0
	} else {
		timeArr[flag] = time.Now().Second()
		flag++
	}
}

// 冒泡排序
func bubbleSort() {
	arr, step, change := getArray(), 0, 0
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			step++
			if arr[i] > arr[j] {
				change++
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	fmt.Printf("冒泡排序: %d 步, 交换 %d 次, 结果：%v \n", step, change, arr)
}

// 选择排序
func selectSort() {
	arr, step, change := getArray(), 0, 0
	length, min := len(arr), 0
	for i := 0; i < length; i++ {
		min = i
		for j := i + 1; j < length; j++ {
			step++
			if arr[min] > arr[j] {
				min = j
			}
		}
		if min != i {
			change++
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
	fmt.Printf("选择排序: %d 步, 交换 %d 次, 结果：%v \n", step, change, arr)
}

// 插入排序
// 1. 从第二个元素开始，每个元素与前面元素比较，每次比较不满足条件就将当前比较值后移
func insertSort() {
	arr, step, change := getArray(), 0, 0
	length := len(arr)
	for i := 1; i < length; i++ {
		val := arr[i]
		compareIndex := i - 1 // 比较位置
		for compareIndex >= 0 && val < arr[compareIndex] {
			change++
			step++
			arr[compareIndex+1] = arr[compareIndex] // 不满足条件就将当前比较值后移
			compareIndex--
		}
		if compareIndex+1 != i {
			change++
			arr[compareIndex+1] = val
		}
	}
	fmt.Printf("插入排序: %d 步, 交换 %d 次, 结果：%v \n", step, change, arr)
}

var step, change int

// 快速排序
// 1. 从第二个元素开始，每个元素与前面元素比较，每次比较不满足条件就将当前比较值后移
func quickSort1(left int, right int, arr *[NUM]int) {

	median := arr[(left+right)/2]
	l, r := left, right

	for r > l {

		// 从左边开始不断寻找大于等于中间值的位置
		for arr[l] <= median {
			l++
		}

		// 从右边开始不断寻找小于等于中间值的位置
		for arr[r] >= median {
			r--
		}

		// 如果左边或者右边已经找到头就退出循环
		if l == right || r == left {
			break
		}

		// 交换位置
		step++
		change++
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}

	if arr[l] == median {
		l++
	}

	if arr[r] == median {
		r--
	}

	fmt.Printf("l=%v; r=%v, median=%v, left=%v, right=%v, arr=%v \n", l, r, median, left, right, arr)

	if r > left {
		// 小于 median 的数据排序
		quickSort1(left, r, arr)
	}

	if l < right {
		// 大于 median 的数据排序
		quickSort1(l, right, arr)
	}
}

func quickSort2(array []int, left, right int) {
	if left >= right {
		return
	}
	pos := partition(array, left, right)
	quickSort2(array, left, pos-1)
	quickSort2(array, pos+1, right)
}

func partition1(array []int, left, right int) int {
	base := array[left]
	for right > left {
		for right > left && array[right] >= base {
			right--
		}
		array[left] = array[right]
		for right > left && array[left] <= base {
			left++
		}
		array[right] = array[left]
	}
	array[right] = base
	return right
}

func partition(arr []int, left, right int) int {
	base := arr[left]
	for right > left {
		// 寻找 left 右边大于基准值的数位置, 并将值赋给 left
		for right > left && arr[right] <= base {
			right--
		}
		change++
		step++
		change++
		step++
		arr[left] = arr[right]
		// 寻找 right 左边小于基准值的数位置, 并将值赋给 right
		for right > left && arr[left] >= base {
			left++
		}
		arr[right] = arr[left]
	}
	// 最后将基准值赋给 right
	arr[right] = base
	return right
}

func main() {

	//fmt.Printf("brefore sort: %v \n", getArray())

	//TODO: 十万条数据测试
	// 冒泡排序
	useTime()
	bubbleSort()
	useTime()

	// 选择排序
	useTime()
	selectSort()
	useTime()

	// 插入排序
	useTime()
	insertSort()
	useTime()

	// 快速排序
	useTime()
	arr := getArray()
	//quickSort1(0, len(arr) - 1, &arr)
	quickSort2(arr[:], 0, len(arr)-1)
	useTime()
	fmt.Printf("快速排序: %d 步, 交换 %d 次, 结果：%v \n", step, change, arr)
}
