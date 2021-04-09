package linkedList

import "errors"

var (
	ErrorInsertBeforeHead = errors.New("数据不能插入头节点之前")
	ErrorNodeNotFound     = errors.New("没有找到节点")
)

type UserDate interface {
	Find(name interface{}) bool
}

const (
	Repeat = iota
	Unique
)

type Link struct {
	head     *doubleNode
	isUnique int
}

func NewLink(userData []UserDate, isUnique int) Link {
	if isUnique != Unique && isUnique != Repeat {
		isUnique = Repeat
	}
	link := Link{&doubleNode{}, isUnique}
	for _, user := range userData {
		link.head.pushNode(&doubleNode{UserDate: user})
	}
	return link
}

func NewNode(userData UserDate) *doubleNode {
	return &doubleNode{UserDate: userData}
}

// 查找节点
func (link *Link) search(name interface{}) (int, *doubleNode, error) {
	temp := link.head.Next
	len := 1
	for {
		if temp != nil && temp.UserDate.Find(name) {
			return len, temp, nil
		}
		if temp.Next != nil {
			len++
			temp = temp.Next
		} else {
			break
		}
	}
	return len, nil, ErrorNodeNotFound
}

// 链表长度
func (link *Link) len() int {
	temp := link.head
	len := 0
	for {
		if temp.Next != nil {
			len++
			temp = temp.Next
		} else {
			break
		}
	}
	return len
}

// 节点
type doubleNode struct {
	Pre      *doubleNode // 上一个数据
	UserDate UserDate    // 用户数据
	Next     *doubleNode // 下一个数据
}

// 节点后面插入数据
func (doubleNode *doubleNode) insertAfterNode(node *doubleNode) {
	if doubleNode.Next != nil {
		temp := doubleNode.Next
		doubleNode.Next = node
		node.Next = temp
		node.Pre = doubleNode
		temp.Pre = node
	} else {
		doubleNode.Next = node
		node.Pre = doubleNode
	}
}

// 节点前面插入数据
func (doubleNode *doubleNode) insertBeforeNode(node *doubleNode) error {
	if doubleNode.Pre != nil {
		temp := doubleNode.Pre
		doubleNode.Pre = node
		node.Pre = temp
		node.Next = doubleNode
		temp.Next = node
		return nil
	} else {
		return ErrorInsertBeforeHead
	}
}

// 删除节点
func (doubleNode *doubleNode) deleteNode() {
	doubleNode.Pre.Next = doubleNode.Next
}

// 尾部追加节点
func (doubleNode *doubleNode) pushNode(node *doubleNode) {
	lastNode := doubleNode.getLastNode()
	lastNode.Next = node
	node.Pre = lastNode
}

// 获取最后一个节点
func (doubleNode *doubleNode) getLastNode() *doubleNode {
	temp := doubleNode
	for {
		if temp.Next == nil {
			return temp
		}
		temp = temp.Next
	}
}
