package stack

import (
	"sync"
)

type BinaryTreeNode struct {
	value     int
	parent    *BinaryTreeNode
	leftNode  *BinaryTreeNode
	rightNode *BinaryTreeNode
}

type BinaryTree struct {
	root      *BinaryTreeNode
	nextIndex int
	mutex     sync.RWMutex
}

func (binaryTree *BinaryTree) Push(element int) {
	binaryTree.mutex.Lock()
	defer binaryTree.mutex.Unlock()

	//fmt.Println("<--", element)

	node := &BinaryTreeNode{element, nil, nil, nil}

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

func (binaryTree *BinaryTree) Pop() (int, bool) {
	binaryTree.mutex.Lock()
	defer binaryTree.mutex.Unlock()

	if binaryTree.root == nil {
		return 0, false
	} else {
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

		//fmt.Println("-->", lastNode.value)

		return lastNode.value, true
	}
}

func (binaryTree *BinaryTree) Print() {
}

func (binaryTree *BinaryTree) findNode(nodeIndex int) *BinaryTreeNode {
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
