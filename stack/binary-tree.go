package stack

import (
	"fmt"
	"sync"
)

type binaryTreeNode struct {
	value     interface{}
	parent    *binaryTreeNode
	leftNode  *binaryTreeNode
	rightNode *binaryTreeNode
}

func (node *binaryTreeNode) print() {
	direction := ""
	var parentValue interface{}

	if node.parent == nil {
		direction = "root"
		parentValue = nil
	} else {
		if node == node.parent.leftNode {
			direction = "L"
		} else if node == node.parent.rightNode {
			direction = "R"
		}
		parentValue = node.parent.value
	}

	fmt.Printf("%4v | %5v | %5v \n", direction, node.value, parentValue)
}

type BinaryTree struct {
	root      *binaryTreeNode
	nextIndex int
	sync.RWMutex
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
	binaryTree.print(binaryTree.root)
}

func (binaryTree *BinaryTree) print(node *binaryTreeNode) {
	if node == nil {
		return
	}

	node.print()

	binaryTree.print(node.leftNode)
	binaryTree.print(node.rightNode)
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

func (binaryTree *BinaryTree) findLevels() int {
	if binaryTree.root == nil {
		return 0
	}

	currentIndex := binaryTree.nextIndex - 1
	start := 0
	end := 0
	levels := 0

	for {
		levels++

		if currentIndex >= start && currentIndex <= end {
			break
		}

		start += start + 1
		end += end + 2
	}

	return levels
}
