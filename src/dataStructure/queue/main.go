package main

import "fmt"

type ArrayQueue struct {
	MaxSize    int   // 队列容量
	Front      int   // 取值位置
	Foot       int   // 加入位置
	ArraySlice []int // 存放队列数据
}

func (aq *ArrayQueue) AddQueue(i int) {
	if aq.Foot == aq.MaxSize-1 {
		fmt.Println("队列已满")
	} else {
		aq.ArraySlice[aq.Foot] = i
		aq.Foot++
		fmt.Println("入列成功: ", i)
	}
}

// 加入环形队列
// 是否已满 : 前后指针距离来判断是否大于容量
// 插入位置 : 加入位置取模最大指针
func (aq *ArrayQueue) AddCircleQueue(i int) {
	// 队列已满
	if aq.Foot-aq.Front > aq.MaxSize-1 {
		fmt.Println("队列已满")
	} else {
		aq.ArraySlice[aq.Foot%aq.MaxSize] = i
		aq.Foot++
		fmt.Println("入列成功: ", i)
	}
}

func (aq *ArrayQueue) PopQueue() {
	if aq.Foot == aq.Front {
		fmt.Println("队列为空")
	} else {
		temp := aq.ArraySlice[aq.Front]
		aq.ArraySlice[aq.Front] = 0
		aq.Front++
		fmt.Println("出列成功: ", temp)
	}
}

// 移出环形队列
// 是否为空 : 前指针是否追上了后指针
// 插入位置 : 取值位置取模最大指针
func (aq *ArrayQueue) PopCircleQueue() {
	if aq.Foot == aq.Front {
		fmt.Println("队列为空")
	} else {
		temp := aq.ArraySlice[aq.Front%aq.MaxSize]
		aq.ArraySlice[aq.Front%aq.MaxSize] = 0
		aq.Front++
		fmt.Println("出列成功: ", temp)
	}
}

func (aq *ArrayQueue) ShowQueue() {
	fmt.Println(aq.ArraySlice)
}

func main() {
	arrayQueue := ArrayQueue{
		MaxSize:    5,
		Front:      0,
		Foot:       0,
		ArraySlice: make([]int, 5),
	}

	var i int
	for {
		fmt.Printf("输入数字: 0(出列) 886(退出) 其他为入列: ")
		fmt.Scanln(&i)
		switch i {
		case 0:
			arrayQueue.PopCircleQueue()
			arrayQueue.ShowQueue()
		case 886:
			fmt.Println("退出成功")
			return
		default:
			arrayQueue.AddCircleQueue(i)
			arrayQueue.ShowQueue()
		}
	}

}
