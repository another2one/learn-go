package main

import "fmt"

// 注意
// 1. 结束条件
// 2. 具体每部操作
// 3. 栈方式调用（先进后出），随调用返回结果给谁

var (
	maze = [5][5]int{
		{0, 1, 0, 0, 0},
		{0, 1, 0, 0, 0},
		{0, 0, 0, 1, 0},
		{0, 1, 1, 0, 0},
		{0, 0, 0, 0, 0},
	}
)

func test(x, y, step int) (res bool) {
	step++
	if x == 4 && y == 4 {
		fmt.Println("走出迷宫了")
		return true
	}
	if maze[x][y] > 0 {
		return false
	} else {
		maze[x][y] = 2
	}

	if y < 4 {
		if test(x, y+1, step) {
			fmt.Println(step)
		}
	}

	if x < 4 {
		if test(x+1, y, step) {
			fmt.Println(step)
		}
	}

	if x > 0 {
		if test(x-1, y, step) {
			fmt.Println(step)
		}
	}
	if y > 0 {
		if test(x, y-1, step) {
			fmt.Println(step)
		}
	}
	return false
}

func main() {

	test(0, 0, 0)
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("%d  ", maze[i][j])
		}
		fmt.Println()
	}

}
