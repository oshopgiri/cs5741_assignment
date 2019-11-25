package stack

import (
	"log"
	"sync"
)

type binaryTreeNode struct {
	value     interface{}
	parent    *binaryTreeNode
	leftNode  *binaryTreeNode
	rightNode *binaryTreeNode
}

type BinaryTree struct {
	root      *binaryTreeNode
	nextIndex int
	sync.RWMutex
}

func (binaryTree *BinaryTree) Init() {}

func (binaryTree *BinaryTree) Push(element interface{}, waitGroup *sync.WaitGroup) {
	if waitGroup != nil {
		defer waitGroup.Done()

		binaryTree.Lock()
		defer binaryTree.Unlock()
	}

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

	log.Println("<--", element)
}

func (binaryTree *BinaryTree) Pop(waitGroup *sync.WaitGroup) (interface{}, bool) {
	if waitGroup != nil {
		defer waitGroup.Done()

		binaryTree.Lock()
		defer binaryTree.Unlock()
	}

	if binaryTree.root == nil {
		if waitGroup != nil {
			waitGroup.Add(1)
			go binaryTree.Pop(waitGroup)
		}

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

		log.Println("-->", lastNode.value)

		return lastNode.value, true
	}
}

func (binaryTree *BinaryTree) Print() {}

func (binaryTreeNode *binaryTreeNode) Print() {}

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
