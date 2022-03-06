package cane

import (
	"fmt"
	"strings"
)

// a prefix tree with path can look like this for example
//                    GET
//                   ┌─────┐
//               ┌───┤     ├─────────┐
//               │   └─────┘         │
//               │                   │
//            ┌──▼──┐           ┌───▼───┐
//    ┌───────┤ doc  │           │contact │
//    │       └──┬───┘           └────────┘
//    │          │
// ┌──▼──┐   ┌───▼───┐
// │sdks  │   │:lang   ├──────┐
// └──────┘   └──┬─────┘      │
//               │            │
//            ┌──▼────┐   ┌──▼────┐
//            │ usage  │   │ spec  │
//            └────────┘   └───────┘
type node struct {
	pattern  string  // full pattern
	part     string  // the part of the path
	children []*node // the available part that can come after this
	isWild   bool    // *
}

func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

func (n *node) matchAllChild(part string) []*node {
	res := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			res = append(res, child)
		}
	}
	return res
}

func (n *node) insert(pattern string, parts []string, height int) {
	// final node in path
	if len(parts) == height {
		if n.pattern != "" {
			// panic if route conflict happens
			panic(fmt.Sprintf("route conflict %s, origin %s", pattern, n.pattern))
		}
		n.pattern = pattern
		return
	}
	currentPart := parts[height]
	//recursively insert parts
	child := n.matchChild(currentPart)
	if child == nil {
		child = &node{
			part:   currentPart,
			isWild: currentPart[0] == ':' || currentPart[0] == '*',
		}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	//found the node at the end of path
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchAllChild(part)

	for _, child := range children {
		node := child.search(parts, height+1)
		if node != nil {
			return node
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

func (n *node) String() string {
	return fmt.Sprintf("node{pattern=%s, part=%s, isWild=%t}", n.pattern, n.part, n.isWild)
}
