// 队列 先进先出的数据结构

// 1. 数组实现
// 环形队列 Foot 和 Front 可以不停往后取值
// 取模判断操作队列的实际位置
// 根据前后距离判断是否为空及已满
type ArrayQueue struct {
	MaxSize int // 队列容量
	Front int // 取值位置
	Foot int // 加入位置
	ArraySlice []int // 存放队列数据
}

// 2. 链表