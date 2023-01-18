package bst

import (
	"github/hsingyingli/data-structure-and-algorithm/util"
)

func height(node *Node) int {
	if node == nil {
		return 0
	}
	return node.size
}

func BuildAVL(root *Node, value int) *Node {
	if root == nil {
		return &Node{value: value, size: 1}
	}

	if value > root.value {
		root.right = BuildAVL(root.right, value)
		root.right.parent = root
	}

	if value < root.value {
		root.left = BuildAVL(root.left, value)
		root.left.parent = root
	}

	root.size = 1 + util.Max(height(root.right), height(root.left))

	return root.balancing()
}

func (node *Node) rotateRight() *Node {
	newRoot := node.left
	children := newRoot.right

	node.left = children
	newRoot.right = node

	newRoot.parent = node.parent
	if children != nil {
		children.parent = node
	}
	node.parent = newRoot

	node.updating()
	return newRoot
}

func (node *Node) rotateLeft() *Node {
	newRoot := node.right
	children := newRoot.left

	node.right = children
	newRoot.left = node

	newRoot.parent = node.parent
	if children != nil {
		children.parent = node
	}

	node.parent = newRoot

	node.updating()

	return newRoot
}

func (node *Node) updating() {
	if node == nil {
		return
	}

	node.size = 1 + util.Max(height(node.right), height(node.left))
	if node.parent != nil {
		node.parent.updating()
	}
}

// skew(node): height(node.left) - height(node.right) => {-1, 0, +1}
func (node *Node) skew() int {
	return height(node.left) - height(node.right)
}

func (node *Node) balancing() *Node {
	skew := node.skew()

	if skew < 2 && skew > -2 {
		return node
	}

	// right left rotate
	if skew == -2 && node.right.skew() == 1 {
		node.right = node.right.rotateRight()
		return node.rotateLeft()
	}

	// left rotate
	if skew == -2 && node.right.skew() == -1 {
		return node.rotateLeft()
	}

	// right rotate
	if skew == 2 && node.left.skew() == 1 {
		return node.rotateRight()
	}

	// left  right rotate
	if skew == 2 && node.left.skew() == -1 {
		node.left = node.left.rotateLeft()
		return node.rotateRight()
	}

	return node
}
