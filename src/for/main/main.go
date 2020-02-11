package main

import(
	"fmt"
)

func main() {

	// 第一种
	for i := 1; i < 3; i++ {
		fmt.Println("666")
	}

	// 第二种
	j := 1
	for j < 3 {
		j++
		fmt.Println("777")
	}

	// 第三种
	s := 1
	for {
		if s > 3 {
			break
		}
		s++
		fmt.Println("888")
	}

	// 第四种
	for index, val := range "sss生导师" {
		fmt.Printf("%d == %c \n", index, val)
	}
	fmt.Printf("%c \n", "666"[0])

	for i := 0; i <= 6; i++ {
		j = 6 - i
		fmt.Printf("%d + %d = 6 \n", i, j)
	}

	// break 默中止当前语句块，可以通过标签中止到指定层
	i1, i2 := 0, 0
	label1:
	for {
		i1++
		if i1 > 3 {
			break
		}
		fmt.Printf("i1 = %d \n", i1)
		i2 = 0
		for {
			i2++
			if i2 == 2 {
				continue
			}
			if i2 > 3 {
				break label1
			}
			fmt.Printf("i2 = %d \n", i2)
		}
	}
}