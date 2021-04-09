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

func (water *water) Find(name interface{}) bool {
	switch name := name.(type) {
	case string:
		return water.Name == name
	}
	return false
}

func TestNew(t *testing.T) {
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

	link = NewLink([]UserDate{Node1, Node2, Node3}, Unique)
	doublLen := link.len()
	if doublLen != 3 {
		log.Fatalf("link len expect to %d but %d", 2, doublLen)
	}
}

func TestDobuleSearch(t *testing.T) {
	index, node, err := link.search("卢俊义")
	if err != nil {
		log.Fatalln("dobule search error: ", err)
	}
	w, ok := node.UserDate.(*water)
	if !ok {
		log.Fatalln("dobule search type error: not water type")
	}
	if w.Name != "卢俊义" || w.Number != 2 || w.NickName != "玉麒麟" {
		log.Fatalf("dobule search elemt error: %+v \n", *w)
	}
	if index != 3 {
		log.Fatalf("dobule search index error: expect %d but %d \n", 3, index)
	}
}

func TestDobuleUpdate(t *testing.T) {
	_, node, err := link.search("宋江")
	if err != nil {
		log.Fatalln("dobule search error: ", err)
	}
	w, ok := node.UserDate.(*water)
	if !ok {
		log.Fatalln("dobule search error: not water type")
	}

	old := w.NickName
	w.NickName = "小人"

	_, node, err = link.search("宋江")
	if err != nil {
		log.Fatalln("dobule search error: ", err)
	}
	w, ok = node.UserDate.(*water)
	if !ok {
		log.Fatalln("dobule search error: not water type")
	}
	if w.NickName != "小人" {
		log.Fatalln("update dobule error")
	} else {
		w.NickName = old
	}
}

func TestInsertBefore(t *testing.T) {
	// "卢俊义"之前插入
	beforeLen := link.len()
	index, node, err := link.search("卢俊义")
	if err != nil {
		log.Fatalln("dobule search error: ", err)
	}
	node.insertBeforeNode(NewNode(&water{
		Name:     "lizhi",
		Number:   4,
		NickName: "666",
	}))
	// 判断插入后长度是否加1
	afterLen := link.len()
	if afterLen != beforeLen+1 {
		log.Fatalf("link len expect to be %d but %d \n", beforeLen+1, afterLen)
	}
	// 判断是否能找到插入元素及index是否为插入位置的index
	indexNew, _, err := link.search("lizhi")
	if err != nil {
		log.Fatalln("dobule insert error: ", err)
	}
	if indexNew != index {
		log.Fatalf("new node index expect to be %d but %d \n", index, indexNew)
	}
}

func TestInsertAfter(t *testing.T) {
	// "卢俊义"之后插入
	beforeLen := link.len()
	index, node, err := link.search("卢俊义")
	if err != nil {
		log.Fatalln("dobule search error: ", err)
	}
	node.insertAfterNode(NewNode(&water{
		Name:     "lipan",
		Number:   7,
		NickName: "777",
	}))
	// 判断插入后长度是否加1
	afterLen := link.len()
	if afterLen != beforeLen+1 {
		log.Fatalf("link len expect to be %d but %d \n", beforeLen+1, afterLen)
	}
	// 判断是否能找到插入元素及index是否为插入位置的index+1
	indexNew, _, err := link.search("lipan")
	if err != nil {
		log.Fatalln("dobule insert error: ", err)
	}
	if indexNew != index+1 {
		log.Fatalf("new node index expect to be %d but %d \n", index+1, indexNew)
	}
}

func TestDelete(t *testing.T) {
	beforeLen := link.len()
	_, node, err := link.search("宋江")
	if err != nil {
		log.Fatalln("dobule search error: ", err)
	}
	node.deleteNode()
	afterLen := link.len()
	if afterLen != beforeLen-1 {
		log.Fatalf("link len expect to %d but %d \n", afterLen, beforeLen-1)
	}
}
