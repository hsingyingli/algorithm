package bst

import "github/hsingyingli/data-structure-and-algorithm/util"

type BST interface {
	height() int
	depth() int

	traversal()

	subtreeFirst() *Node
	subtreeAfter() *Node

	successor() *Node
	predecessor() *Node

	subtreeInsertAfter(new *Node)
	subtreeInsertBefore(new *Node)

	subtreeDelete()
}

type Node struct {
	left   *Node
	right  *Node
	parent *Node
	value  int
}

func BuildBST(root *Node, value int) {
	if root.value > value {
		if root.left != nil {
			BuildBST(root.left, value)
		} else {
			root.left = &Node{
				parent: root,
				value:  value,
			}
		}
	}

	if root.value <= value {
		if root.right != nil {
			BuildBST(root.right, value)
		} else {
			root.right = &Node{
				parent: root,
				value:  value,
			}
		}
	}
}

// height return the height of the node and it's descendants
func (node *Node) height() int {
	height := 0
	left := 0
	right := 0
	if node.left != nil || node.right != nil {
		height++
	}

	if node.left != nil {
		left = node.left.height()
	}
	if node.right != nil {
		right = node.right.height()
	}
	return height + util.Max(left, right)
}

// depth number of edge in the path from node to root
func (node Node) depth() int {
	depth := 0

	if node.parent != nil {
		depth++
		depth += node.parent.depth()
	}

	return depth
}

// traversal return node.value in traversal order
func (node *Node) traversal() []*Node {
	order := []*Node{}
	if node.left != nil {
		order = append(node.left.traversal(), order...)
	}
	order = append(order, node)
	if node.right != nil {
		order = append(order, node.right.traversal()...)
	}
	return order
}

// subtreeFirst which node come first in traversal order (left most)
func (node *Node) subtreeFirst() *Node {
	if node.left == nil {
		return node
	}
	return node.left.subtreeFirst()
}

// subtreeAfter which node come after in traversal order (right most)
func (node *Node) subtreeAfter() *Node {
	if node.right == nil {
		return node
	}
	return node.right.subtreeAfter()
}

// successor return universal node come after in traversal order
func (node *Node) successor() *Node {
	if node.right != nil {
		return node.right.subtreeFirst()
	}

	parent := node.parent
	child := node

	for parent != nil && child != parent.left {
		child = parent
		parent = child.parent
	}
	return parent
}

// predecessor return universal node come before in traversal order
func (node *Node) predecessor() *Node {
	if node.left != nil {
		return node.left.subtreeAfter()
	}

	parent := node.parent
	child := node

	for parent != nil && child != parent.right {
		child = parent
		parent = child.parent
	}
	return parent
}

func (node *Node) subtreeInsertAfter(new *Node) {
	if node.right == nil {
		node.right = new
		new.parent = node
		return
	}

	successor := node.successor()
	successor.left = new
	new.parent = successor
}

func (node *Node) subtreeInsertBefore(new *Node) {
	if node.left == nil {
		node.left = new
		new.parent = node
		return
	}

	predecessor := node.predecessor()
	predecessor.right = new
	new.parent = predecessor
}

func (node *Node) subtreeDelete() {
	// if node is leaf
	if node.left == nil && node.right == nil && node.parent != nil {
		if node.parent.left == node {
			node.parent.left = nil
			node.parent = nil
		} else {
			node.parent.right = nil
			node.parent = nil
		}
		return
	}

	if node.left != nil {
		predecessor := node.predecessor()
		node.value = predecessor.value
		predecessor.subtreeDelete()
		return
	}

	if node.right != nil {
		successor := node.successor()
		node.value = successor.value
		successor.subtreeDelete()
	}
}
