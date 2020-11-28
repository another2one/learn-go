package main

import (
	"errors"
	"fmt"
)

type WaterNode struct {
	Pre      *WaterNode
	Name     string
	Number   int
	NickName string
	Next     *WaterNode
}

var waterHeadNode = &WaterNode{}

// wn 之后加入节点
func (wn *WaterNode) InsertAfterNode(node *WaterNode) {
	if wn.Next != nil {
		temp := wn.Next
		wn.Next = node
		node.Next = temp
		node.Pre = wn
		temp.Pre = node
	} else {
		wn.Next = node
		node.Pre = wn
	}
}

// wn 之前加入节点
func (wn *WaterNode) InsertBeforeNode(node *WaterNode) {
	if wn.Pre != nil {
		temp := wn.Pre
		wn.Pre = node
		node.Pre = temp
		node.Next = wn
		temp.Next = node
	} else {
		fmt.Println("不能插入头节点之前")
	}
}

// 尾部追加节点
func (wn *WaterNode) pushNode(node *WaterNode) {
	lastNode := wn.getLastNode()
	lastNode.Next = node
	node.Pre = lastNode
}

func (wn *WaterNode) getLastNode() *WaterNode {
	temp := wn
	for {
		if temp.Next == nil {
			return temp
		}
		temp = temp.Next
	}
}

func (wn *WaterNode) lens() int {
	tempNext := wn.Next
	tempPre := wn.Pre
	var len = 1
	for {
		if tempNext.Next != nil {
			tempNext = tempNext.Next
			len++
		}
		if tempPre.Pre != nil {
			tempPre = tempPre.Pre
			len++
		}
	}
	return len
}

func (wn *WaterNode) insertByNumber(node *WaterNode) {
	temp := waterHeadNode
	for {
		if temp.Next == nil || temp.Next.Number >= node.Number {
			temp.InsertAfterNode(node)
			break
		}
		temp = temp.Next
	}
}

func (wn *WaterNode) insertByNumberWithoutRepeat(node *WaterNode) {
	temp := waterHeadNode
	for {
		if temp.Next == nil || temp.Next.Number > node.Number {
			temp.InsertAfterNode(node)
			break
		} else if temp.Next.Number == node.Number {
			fmt.Println("sorry! 有钱真的可以为所欲为 ！！！但是不能重复插入")
		}
		temp = temp.Next
	}
}

func (wn *WaterNode) deleteSelf() {

	if wn.Pre == nil {
		// 头节点不能删除
		return
	}

	if wn.Next != nil {
		wn.Pre.Next = wn.Next
		wn.Next.Pre = wn.Pre
	} else {
		// 尾节点不能删除
		wn.Pre.Next = nil
	}

}

func (wn *WaterNode) searchNode(name string) (resNode *WaterNode, err error) {
	if wn.Name == name {
		return wn, nil
	}
	tempNext := wn.Next
	tempPre := wn.Pre
	for {
		if tempNext.Name == name {
			resNode = tempNext
		} else if tempNext.Next != nil {
			tempNext = tempNext.Next
		}
		if tempPre.Name == name {
			resNode = tempPre
		} else if tempPre.Pre != nil {
			tempPre = tempPre.Pre
		}
	}
	return resNode, errors.New("not found")
}

func (wn *WaterNode) showNode() {
	temp := waterHeadNode
	for {
		if temp.Next == nil {
			break
		}
		temp = temp.Next
		fmt.Println(temp)
	}
}

func main() {

	Node1 := &WaterNode{
		Name:     "宋江",
		Number:   1,
		NickName: "及时雨",
		Next:     nil,
	}
	waterHeadNode.pushNode(Node1)
	Node1.InsertAfterNode(&WaterNode{
		Name:     "吴用",
		Number:   3,
		NickName: "智多星",
		Next:     nil,
	})

	node3 := &WaterNode{
		Name:     "卢俊义",
		Number:   2,
		NickName: "玉麒麟",
		Next:     nil,
	}
	waterHeadNode.insertByNumber(node3)

	waterHeadNode.showNode()

	node3.deleteSelf()

	waterHeadNode.showNode()
}
