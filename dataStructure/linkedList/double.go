package linkedList

import "errors"

var (
	ErrorInsertBeforeHead = errors.New("数据不能插入头节点之前")
	ErrorNodeNotFound     = errors.New("没有找到节点")
)

type UserDate interface {
	find(v interface{}) bool
}

type Link struct {
	head *doubleNode
}

func NewLink(userData []*UserDate) Link {
	link := Link{&doubleNode{}}
	for _, user := range userData {
		link.head.pushNode(&doubleNode{UserDate: user})
	}
	return link
}

// 查找节点
func (link *Link) search(v interface{}) (*doubleNode, error) {
	temp := link.head.Next
	for {
		if UserDate.find(v) {
			return temp, nil
		}
		if temp.Next == nil {
			temp = temp.Next
		} else {
			break
		}
	}
	return nil, ErrorNodeNotFound
}

// 节点
type doubleNode struct {
	Pre      *doubleNode // 上一个数据
	UserDate *UserDate   // 用户数据
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
