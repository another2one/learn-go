package main

import "fmt"

// 单链表
type HeroNode struct {
	Name     string
	Number   int
	NickName string
	Next     *HeroNode
}

var headNode = &HeroNode{}

// hn 之后加入节点
func (hn *HeroNode) InsertAfterNode(node *HeroNode) {
	if hn.Next != nil {
		temp := hn.Next
		hn.Next = node
		node.Next = temp
	} else {
		hn.Next = node
	}
}

// hn 之前加入节点
func (hn *HeroNode) InsertBeforeNode(node *HeroNode) {
	preNode := hn.getPreNode(headNode)
	if preNode == nil {
		return
	}
	preNode.Next = node
	node.Next = hn
}

// 尾部追加节点
func (hn *HeroNode) pushNode(node *HeroNode) {
	hn.getLastNode(hn).Next = node
}

func (hn *HeroNode) getPreNode(startNode *HeroNode) *HeroNode {
	if startNode.Next == hn {
		return startNode
	} else if startNode.Next == nil {
		return nil
	} else {
		return hn.getPreNode(startNode.Next)
	}
}

func (hn *HeroNode) getLastNode(startNode *HeroNode) *HeroNode {
	if startNode.Next == nil {
		return startNode
	} else {
		return hn.getLastNode(startNode.Next)
	}
}

func (hn *HeroNode) insertByNumber(node *HeroNode) {
	temp := headNode
	for {
		if temp.Next == nil || temp.Next.Number > node.Number {
			temp.InsertAfterNode(node)
			break
		} else if temp.Next.Number == node.Number {
			fmt.Println("sorry! 有钱真的可以为所欲为 ！！！")
		}
		temp = temp.Next
	}
}

func (hn *HeroNode) deleteNode() {
	preNode := hn.getPreNode(headNode)
	if preNode == nil {
		return
	}
	if hn.Next == nil {
		preNode.Next = nil
	} else {
		preNode.Next = hn.Next
	}

}

func (hn *HeroNode) searchNode(startNode *HeroNode, name string) *HeroNode {
	if startNode.Name == name {
		return startNode
	} else {
		return hn.searchNode(startNode.Next, name)
	}
}

func (hn *HeroNode) showNode() {
	temp := headNode
	for {
		if temp.Next == nil {
			break
		}
		temp = temp.Next
		fmt.Println(temp)
	}
}

func main() {

	Node1 := &HeroNode{
		Name:     "宋江",
		Number:   1,
		NickName: "及时雨",
		Next:     nil,
	}
	headNode.pushNode(Node1)
	Node1.InsertAfterNode(&HeroNode{
		Name:     "吴用",
		Number:   3,
		NickName: "智多星",
		Next:     nil,
	})

	node3 := &HeroNode{
		Name:     "卢俊义",
		Number:   2,
		NickName: "玉麒麟",
		Next:     nil,
	}
	headNode.insertByNumber(node3)

	headNode.showNode()

	node3.deleteNode()

	headNode.showNode()
}
