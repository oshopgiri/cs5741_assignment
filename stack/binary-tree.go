package stack

import (
	"fmt"
	"strings"
)

type binaryTreeNode struct {
	value     interface{}
	parent    *binaryTreeNode
	leftNode  *binaryTreeNode
	rightNode *binaryTreeNode
}

func (node *binaryTreeNode) print(level int) {
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

type BinaryTree struct {
	root      *binaryTreeNode
	nextIndex int
}

func (binaryTree *BinaryTree) Init() {}

func (binaryTree *BinaryTree) Push(element interface{}) {
	node := &binaryTreeNode{element, nil, nil, nil}

	if binaryTree.root == nil {
		binaryTree.root = node
	} else {
		parent := binaryTree.findNode((binaryTree.nextIndex - 1) / 2)
		if binaryTree.nextIndex%2 == 0 {
			parent.rightNode = node
		} else {
			parent.leftNode = node
		}
		node.parent = parent
	}

	binaryTree.nextIndex++
}

func (binaryTree *BinaryTree) Pop() (interface{}, bool) {
	if binaryTree.root == nil {
		return nil, false
	}

	lastNodeIndex := binaryTree.nextIndex - 1
	lastNode := binaryTree.findNode(lastNodeIndex)

	if lastNodeIndex == 0 {
		binaryTree.root = nil
	} else {
		if lastNodeIndex%2 == 0 {
			lastNode.parent.rightNode = nil
		} else {
			lastNode.parent.leftNode = nil
		}
	}

	binaryTree.nextIndex--

	return lastNode.value, true
}

func (binaryTree *BinaryTree) Print() {
	fmt.Println(strings.Repeat("-", 50))
	binaryTree.root.print(0)
	fmt.Println(strings.Repeat("-", 50))
}

func (binaryTree *BinaryTree) findNode(nodeIndex int) *binaryTreeNode {
	if nodeIndex == 0 {
		return binaryTree.root
	} else if nodeIndex > 0 {
		pathToRoot := binaryTree.findPathToRoot(nodeIndex)

		currentNode := binaryTree.root
		for i := len(pathToRoot) - 1; i >= 0; i-- {
			if pathToRoot[i]%2 == 0 {
				currentNode = currentNode.rightNode
			} else {
				currentNode = currentNode.leftNode
			}
		}

		return currentNode
	} else {
		return nil
	}
}

func (binaryTree *BinaryTree) findPathToRoot(nodeIndex int) []int {
	pathToRoot := []int{nodeIndex}
	currentNodeIndex := nodeIndex

	for {
		parentIndex := (currentNodeIndex - 1) / 2
		if parentIndex == 0 {
			break
		}
		pathToRoot = append(pathToRoot, parentIndex)
		currentNodeIndex = parentIndex
	}

	return pathToRoot
}
