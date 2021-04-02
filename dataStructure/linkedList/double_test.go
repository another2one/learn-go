package linkedList

import (
	"log"
	"testing"
)

type water struct {
	Name     string
	Number   int64
	NickName string
}

var link Link

func (water *water) find(name interface{}) bool {
	switch name.(type) {
	case string:
		return water.Name == name
	}
	log.Fatal("user find 参数错误")
	return false
}

func setup() {
	var Node1, Node2, Node3 UserDate

	Node1 = &water{
		Name:     "宋江",
		Number:   1,
		NickName: "及时雨",
	}
	Node2 = &water{
		Name:     "吴用",
		Number:   3,
		NickName: "智多星",
	}
	Node3 = &water{
		Name:     "卢俊义",
		Number:   2,
		NickName: "玉麒麟",
	}

	link = NewLink([]*UserDate{&Node1, &Node2, &Node3})
	log.Fatal()
}

func TestDobule(t *testing.T) {

}
