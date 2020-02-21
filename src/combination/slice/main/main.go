package main

import(
	"fmt"
)

func fb (fbSlice []int) func () {
	return func(){
		fblen := len(fbSlice)
		fbSlice = append(fbSlice, fbSlice[fblen-1] + fbSlice[fblen-2])
		fmt.Println("fbSlice = ", fbSlice)
	}
}

// 本质就是slice结构体

func main() {

	// make(type, len, [cap])
	makeSlice := make([]int, 2, 5)
	fmt.Printf("makeSlice = %v, type = %T, 容量 = %d \n", makeSlice, makeSlice, cap(makeSlice))

	// 直接定义
	makeSlice1 := []int{1, 3, 5}
	fmt.Printf("makeSlice1 = %v, type = %T, 容量 = %d \n", makeSlice1, makeSlice1, cap(makeSlice1))


	// 数组引用
	var students1 = [5]int{1, 2, 3, 4, 5}
	students2 := students1[1:3]
	// append扩容
	students2[1] = 44
	fmt.Printf("students2 = %v, type = %T, 容量 = %d \n", students2, students2, cap(students2))
	students2 = append(students2, 200, 300) // 追加数量超过cap容量后地址变量
	fmt.Printf("students2 = %v, type = %T, 容量 = %d \n", students2, students2, cap(students2))
	fmt.Printf("students1 = %v \n", students1)
	fmt.Printf("students1 的地址是 %p, students1[1] 的地址是 %p, students2 的地址是 %p, students2[0] 的地址是 %p \n", &students1, &students1[1], &students2, &students2[0])

	// 拷贝
	fmt.Printf("makeSlice = %v, students2 = %v \n", makeSlice, students2)
	copy(makeSlice, students2)
	fmt.Printf("makeSlice = %v, students2 = %v \n", makeSlice, students2)

	fbFunc := fb([]int{1, 1})
	fbFunc()
}