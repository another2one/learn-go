package linkedList

import (
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
func (wn *WaterNode) PushNode(node *WaterNode) {
	lastNode := wn.GetLastNode()
	lastNode.Next = node
	node.Pre = lastNode
}

func (wn *WaterNode) GetLastNode() *WaterNode {
	temp := wn
	for {
		if temp.Next == nil {
			return temp
		}
		temp = temp.Next
	}
}

func (wn *WaterNode) Lens() int {
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
}

func (wn *WaterNode) InsertByNumber(node *WaterNode) {
	temp := waterHeadNode
	for {
		if temp.Next == nil || temp.Next.Number >= node.Number {
			temp.InsertAfterNode(node)
			break
		}
		temp = temp.Next
	}
}

func (wn *WaterNode) InsertByNumberWithoutRepeat(node *WaterNode) {
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

func (wn *WaterNode) DeleteSelf() {

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

func (wn *WaterNode) SearchNode(name string) (resNode *WaterNode, err error) {
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
}

func (wn *WaterNode) ShowNode() {
	temp := waterHeadNode
	for {
		if temp.Next == nil {
			break
		}
		temp = temp.Next
		fmt.Println(temp)
	}
}
