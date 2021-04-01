package linkedList

import "fmt"

type SingleCircleNode struct {
	Name     string
	Number   int
	NickName string
	Next     *SingleCircleNode
}

// hn 之后加入节点
func (scn *SingleCircleNode) InsertAfterNode(node *SingleCircleNode) {
	if scn != node {
		scn.Next = node
	}
}

// hn 之前加入节点
func (scn *SingleCircleNode) InsertBeforeNode(node *SingleCircleNode) {
	preNode := scn.GetPreNode()
	if preNode == nil {
		return
	}
	preNode.Next = node
	node.Next = scn
}

func (scn *SingleCircleNode) GetPreNode() *SingleCircleNode {
	temp := scn.Next
	for {
		if temp == scn {
			return temp
		}
		temp = temp.Next
	}
}

func (scn *SingleCircleNode) InsertByNumber(node *SingleCircleNode) {
	temp := scn
	for {
		if temp.Next.Number > node.Number {
			temp.InsertAfterNode(node)
			break
		} else if temp.Next.Number == node.Number {
			fmt.Println("sorry! 有钱真的可以为所欲为 ！！！")
		}
		temp = temp.Next
	}
}

func (scn *SingleCircleNode) DeleteSelf() {
	preNode := scn.GetPreNode()
	if preNode == nil {
		return
	}
	preNode.Next = scn.Next
}

func (scn *SingleCircleNode) SearchNode(name string) *SingleCircleNode {
	if scn.Name == name {
		return scn
	}
	temp := scn.Next
	for {
		if temp == scn {
			return nil
		}
		if temp.Name == name {
			return temp
		}
		temp = temp.Next
	}
}

func (scn *SingleCircleNode) ShowNode() {
	fmt.Println("show node")
	temp := scn.Next
	for {
		fmt.Println(temp)
		if temp == scn {
			break
		}
		temp = temp.Next
	}
}

func NewCircleNode(params ...*SingleCircleNode) *SingleCircleNode {
	for index, val := range params {
		if index == len(params)-1 {
			val.Next = params[0]
		} else {
			val.Next = params[index+1]
		}
	}
	return params[0]
}

func RunSingleCircle() {

	Node1 := &SingleCircleNode{
		Name:     "宋江",
		Number:   1,
		NickName: "及时雨",
		Next:     nil,
	}
	Node2 := &SingleCircleNode{
		Name:     "吴用",
		Number:   3,
		NickName: "智多星",
		Next:     Node1,
	}
	Node := NewCircleNode(Node1, Node2)
	Node.ShowNode()

	node3 := &SingleCircleNode{
		Name:     "卢俊义",
		Number:   2,
		NickName: "玉麒麟",
		Next:     nil,
	}
	Node.InsertByNumber(node3)

	Node.ShowNode()

	node3.DeleteSelf()

	Node.ShowNode()
}
