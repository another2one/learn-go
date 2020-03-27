package main

import (
	"dataStructure/test/bankqueue/utils"
	"fmt"
)

// 实现银行排队叫号， n 个窗口， m 个人排队， 每个窗口完成后就处理队列其他数据

func main() {

	windowNum, personNum := 4, 9

	fmt.Printf("\n 数组队列实现 \n")
	utils.ArrayQueueRun(windowNum, personNum)

	fmt.Printf("\n 链表实现 \n")
	utils.LinkListRun(windowNum, personNum)

	fmt.Printf("\n 通道实现 \n")
	utils.ChanRun(windowNum, personNum)

}
