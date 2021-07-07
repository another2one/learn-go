package gee

import (
	"strings"
)

type node struct {
	pattern  string  // 待匹配路由，例如 /p/:lang
	part     string  // 路由中的一部分，例如 :lang
	children []*node // 子节点，例如 [doc, tutorial, intro]
	isWild   bool    // 是否精确匹配，part 含有 : 或 * 时为true
}

// 第一个匹配成功的节点，用于插入
func (n *node) matchChild(part string) *node {
	if n.children != nil {
		for _, node := range n.children {
			if node.part == part || node.isWild {
				return node
			}
		}
	}

	return nil
}

// 所有匹配成功的节点，用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	if n.children != nil {
		for _, node := range n.children {
			if node.part == part || node.isWild {
				nodes = append(nodes, node)
			}
		}
	}
	return nodes
}

// lang zh
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}
	nowPattern := parts[height]
	child := n.matchChild(nowPattern)
	if child == nil {
		child = &node{
			part:   nowPattern,
			isWild: nowPattern[0] == ':' || nowPattern[0] == '*',
		}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if height == len(parts) || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}
	pattern := parts[height]
	children := n.matchChildren(pattern)
	for _, child := range children {
		return child.search(parts, height+1)
	}
	return nil
}
