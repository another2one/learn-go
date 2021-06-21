package main

import "fmt"

var N = 20

// 协程交替执行,使其能顺序输出1-20的自然数
func main() {
	c1 := make(chan int) // 有缓存通道不会在超出最大缓存数据是才会写入阻塞，写满时不会
	exitChan := make(chan int)
	// defer close(exitChan)
	defer close(c1)

	go func() {
		if N%2 == 1 { // 如果没有取出就提前关闭了会导致另一个读取异常 所以必须读取后再关闭
			defer close(exitChan)
		}
		for i := 1; i <= N; i++ {
			if i%2 == 0 {
				c1 <- i
			} else {
				fmt.Println("g1:", <-c1)
			}
		}

	}()
	go func() {
		if N%2 == 0 {
			defer close(exitChan)
		}
		for i := 1; i <= N; i++ {
			if i%2 == 0 {
				fmt.Println("g2:", <-c1)
			} else {
				c1 <- i
			}
		}
	}()
	s := <-exitChan // 协程退出后，如果 exitChan 没有关闭的情况下，没有写入将会抛错
	fmt.Println(s)
}
