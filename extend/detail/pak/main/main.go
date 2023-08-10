package main

import (
	"fmt"
	"learn-go/extend/detail/pak/b"
	"learn-go/extend/detail/pak/c"
)

var a = 1

func init() {
	fmt.Println("main.init ...")
}

// question:
//  1. init引入顺序 a 引入 b,c, b引入c abc init顺序
//     先按文件排序引入包 -> 初始化变量 -> init 递归执行上述操作 同一文件只初始化一次
//  2. 包变量定义和init的顺序，由什么决定
//  3. 作用域注意事项：
//     3.1 可重复声明最近层向外找
//     3.2 if 后面表达式也属于 if 作用域
func main() {

	// 1 2
	fmt.Println(b.C)
	fmt.Println(c.C)

	// 3
	if s := 2; s > 1 {
		a := "sss"
		fmt.Println(a)
	}
	fmt.Println("a = ", a)
	// fmt.Println("s = ", s) error
}
