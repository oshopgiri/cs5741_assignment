package avl_tree

import (
	"fmt"
	"math"
	"strings"
)

type avlTreeNode struct {
	value     int
	parent    *avlTreeNode
	leftNode  *avlTreeNode
	rightNode *avlTreeNode
	height    int
	count     int
}

func initTreeNode(element int) *avlTreeNode {
	return &avlTreeNode{element, nil, nil, nil, 0, 0}
}

func (node *avlTreeNode) isBalanced() bool {
	leftHeight, rightHeight := 0, 0

	if node.leftNode != nil {
		leftHeight = node.leftNode.height
	}

	if node.rightNode != nil {
		rightHeight = node.rightNode.height
	}

	return math.Abs(float64(leftHeight)-float64(rightHeight)) < 2
}

func (node *avlTreeNode) print(level int) {
	format := ""

	for i := 0; i < level; i++ {
		format += strings.Repeat(" ", 10)
	}

	format += "---[ "
	level++

	if node.rightNode != nil {
		node.rightNode.print(level)
	}

	fmt.Printf(format+"%d\n", node.value)

	if node.leftNode != nil {
		node.leftNode.print(level)
	}
}

func (node *avlTreeNode) updateHeight() {
	left, right := 0, 0

	if node.leftNode != nil {
		left = node.leftNode.height
	}

	if node.rightNode != nil {
		right = node.rightNode.height
	}

	node.height = int(math.Max(float64(left), float64(right))) + 1
}

type AVLTree struct {
	root *avlTreeNode
}

func (avlTree *AVLTree) Init() {}

func (avlTree *AVLTree) Insert(element int) {
	if avlTree.root == nil {
		avlTree.root = initTreeNode(element)
		avlTree.root.updateHeight()
	} else {
		avlTree.insert(avlTree.root, nil, element)
	}
}

func (avlTree *AVLTree) Delete(element int) (int, bool) {


	return 0, false
}

func (avlTree *AVLTree) Print() {
	fmt.Println(strings.Repeat("-", 50))
	avlTree.root.print(0)
	fmt.Println(strings.Repeat("-", 50))
}

func (avlTree *AVLTree) insert(node *avlTreeNode, parentNode *avlTreeNode, element int) {
	if node == nil {
		node = initTreeNode(element)
		node.parent = parentNode
		if element < parentNode.value {
			parentNode.leftNode = node
		} else {
			parentNode.rightNode = node
		}
		node.updateHeight()
	} else if element == node.value {
		node.count++
	} else if element < node.value {
		avlTree.insert(node.leftNode, node, element)
		if !node.isBalanced() {
			if element < node.leftNode.value {
				avlTree.rotateRight(node)
			} else {
				avlTree.rotateLeftRight(node)
			}
		}
	} else if element > node.value {
		avlTree.insert(node.rightNode, node, element)
		if !node.isBalanced() {
			if element > node.rightNode.value {
				avlTree.rotateLeft(node)
			} else {
				avlTree.rotateRightLeft(node)
			}
		}
	}

	node.updateHeight()
}

func (avlTree *AVLTree) rotateLeft(node *avlTreeNode) {
	rightNode := node.rightNode
	node.rightNode = rightNode.leftNode
	if node.rightNode != nil {
		node.rightNode.parent = node
	}
	rightNode.leftNode = node
	rightNode.parent = node.parent
	if node.parent != nil {
		if rightNode.value > node.parent.value {
			node.parent.rightNode = rightNode
		} else {
			node.parent.leftNode = rightNode
		}
	} else {
		avlTree.root = rightNode
	}
	node.parent = rightNode

	node.updateHeight()
	rightNode.updateHeight()
}

func (avlTree *AVLTree) rotateLeftRight(node *avlTreeNode) {
	avlTree.rotateLeft(node.leftNode)
	avlTree.rotateRight(node)
}

func (avlTree *AVLTree) rotateRight(node *avlTreeNode) {
	leftNode := node.leftNode
	node.leftNode = leftNode.rightNode
	if node.leftNode != nil {
		node.leftNode.parent = node
	}
	leftNode.rightNode = node
	leftNode.parent = node.parent
	if node.parent != nil {
		if leftNode.value > node.parent.value {
			node.parent.rightNode = leftNode
		} else {
			node.parent.leftNode = leftNode
		}
	} else {
		avlTree.root = leftNode
	}
	node.parent = leftNode

	node.updateHeight()
	leftNode.updateHeight()
}

func (avlTree *AVLTree) rotateRightLeft(node *avlTreeNode) {
	avlTree.rotateRight(node.rightNode)
	avlTree.rotateLeft(node)
}
