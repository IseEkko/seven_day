package gee

import (
	"fmt"
	"strings"
)

type node struct {
	pattern  string
	part     string
	children []*node
	isWild   bool
}

func (n *node) String() string {
	return fmt.Sprintf("node{pattern=%s, part=%s, isWild=%t}", n.pattern, n.part, n.isWild)
}

func (n *node) insert(pattern string, parts []string, height int) {

	//查看是不是最后一个，是的话就返回
	if len(parts) == height {
		n.pattern = pattern
		return
	}
	//取出路径中的一部分
	part := parts[height]
	//查看现在的树中是不是有这个节点，没有的话就进行创建
	child := n.matchChild(part)
	//下面是创建的流程。也就是创建
	if child == nil {
		//如果是动态的就会设置为true
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

//路由的查找
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}

func (n *node) travel(list *([]*node)) {
	if n.pattern != "" {
		*list = append(*list, n)
	}
	for _, child := range n.children {
		child.travel(list)
	}
}

//创建节点插入的时候的判断函数
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

//查找路径
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}
