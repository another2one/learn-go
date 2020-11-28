package main

import "fmt"

// 根节点  叶节点  子节点  父节点
// 前序遍历: 根节点开始，左到右子节点遍历
// 中序遍历:
// 后序遍历:

type HeroNode struct {
	Name                byte
	LeftNode, RightNode *HeroNode
}

func getHeroNode(a byte) *HeroNode {
	return &HeroNode{
		Name:      a,
		LeftNode:  nil,
		RightNode: nil,
	}
}

func getAllNode(start, end byte) []*HeroNode {
	res := make([]*HeroNode, end-start+1)
	index := 0
	for i := start; i <= end; i++ {
		res[index] = getHeroNode(i)
		index++
	}
	return res
}

// 前序遍历 a b d h i e j c f g
func PreOrderTraversal(a *HeroNode) {
	fmt.Printf("%c ", a.Name)
	if a.LeftNode != nil {
		PreOrderTraversal(a.LeftNode)
	}
	if a.RightNode != nil {
		PreOrderTraversal(a.RightNode)
	}
}

// 中序遍历 h d i b j e a f c g
func MidOrderTraversal(a *HeroNode) {
	if a.LeftNode != nil {
		MidOrderTraversal(a.LeftNode)
	}
	fmt.Printf("%c ", a.Name)
	if a.RightNode != nil {
		MidOrderTraversal(a.RightNode)
	}
}

// 后序遍历 h i d j e b f g c a
func AfterOrderTraversal(a *HeroNode) {
	if a.LeftNode != nil {
		AfterOrderTraversal(a.LeftNode)
	}
	if a.RightNode != nil {
		AfterOrderTraversal(a.RightNode)
	}
	fmt.Printf("%c ", a.Name)
}

func main() {
	res := getAllNode('a', 'j')
	a, b, c, d, e, f, g, h, i, j := res[0], res[1], res[2], res[3], res[4], res[5], res[6], res[7], res[8], res[9]
	d.LeftNode = h
	d.RightNode = i
	e.LeftNode = j
	b.LeftNode = d
	b.RightNode = e
	c.LeftNode = f
	c.RightNode = g
	a.LeftNode = b
	a.RightNode = c
	fmt.Println("前序排列：")
	PreOrderTraversal(a)
	fmt.Println("\n中序排列：")
	MidOrderTraversal(a)
	fmt.Println("\n后序排列：")
	AfterOrderTraversal(a)
}
