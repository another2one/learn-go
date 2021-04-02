package linkedList

import "testing"

func TestSingleCircle(t *testing.T) {

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
