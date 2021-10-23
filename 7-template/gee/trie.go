package gee

import (
	"strings"
)

// 路由的作用注册与匹配
// 开发服务时，注册路由规则，映射handler；访问时，匹配路由规则，查找到对应的handler。
// Trie 树需要支持节点的插入与查询。

type node struct {
	pattern  string  // 待匹配路由，例如 /p/:lang
	part     string  // 路由中的一部分，例如 :lang
	children []*node // 子节点，例如 [doc, tutorial, intro]
	isWild   bool    // 是否精确匹配，part 含有 : 或 * 时为true
}

// 第一个匹配成功的节点，用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}

	return nil
}

// 所有匹配成功的节点，用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}

	return nodes
}

// 插入功能很简单，递归查找每一层的节点，如果没有匹配到当前part的节点，则新建一个
func (n *node) insert(pattern string, parts []string, height int) {

	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	// 该层匹配不到，则新增
	if child == nil {
		child = &node{
			part:   part,
			isWild: part[0] == ':' || part[0] == '*',
		}
		n.children = append(n.children, child)
	}
	// 往下一层插入
	child.insert(pattern, parts, height+1)
}

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
