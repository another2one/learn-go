package linkedList

import "testing"

func TestDobule(t *testing.T) {

	Node1 := &WaterNode{
		Name:     "宋江",
		Number:   1,
		NickName: "及时雨",
		Next:     nil,
	}
	waterHeadNode.PushNode(Node1)
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
	waterHeadNode.InsertByNumber(node3)

	waterHeadNode.ShowNode()

	node3.DeleteSelf()

	waterHeadNode.ShowNode()
}
