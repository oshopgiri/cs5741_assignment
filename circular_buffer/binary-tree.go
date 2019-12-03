package circular_buffer

import (
	"fmt"
)

type binaryTreeNode struct {
	value     interface{}
	parent    *binaryTreeNode
	leftNode  *binaryTreeNode
	rightNode *binaryTreeNode
}

func (node *binaryTreeNode) print(readPointerNode, writePointerNode *binaryTreeNode) {
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

	fmt.Printf("%4v | %5v | %5v", direction, node.value, parentValue)
	if node == readPointerNode {
		fmt.Print(" <-- readPointer")
	}
	if node == writePointerNode {
		fmt.Print(" <-- writePointer")
	}

	fmt.Println()
}

type BinaryTree struct {
	root         *binaryTreeNode
	readPointer  int
	writePointer int
	size         int
}

func (binaryTree *BinaryTree) Init(size int) {
	for i := 0; i < size; i++ {
		node := &binaryTreeNode{nil, nil, nil, nil}

		if binaryTree.root == nil {
			binaryTree.root = node
		} else {
			parent := binaryTree.findNode((i - 1) / 2)
			if i%2 == 0 {
				parent.rightNode = node
			} else {
				parent.leftNode = node
			}
			node.parent = parent
		}
	}

	binaryTree.readPointer = 0
	binaryTree.writePointer = 0
	binaryTree.size = size
}

func (binaryTree *BinaryTree) Write(element interface{}) bool {
	writeNode := binaryTree.findNode(binaryTree.writePointer)
	if writeNode.value == nil {
		writeNode.value = element
		binaryTree.writePointer = (binaryTree.writePointer + 1) % binaryTree.size

		fmt.Println("WRITE", element)
		binaryTree.Print()
		fmt.Println()

		return true
	}

	return false
}

func (binaryTree *BinaryTree) Read() (interface{}, bool) {
	readNode := binaryTree.findNode(binaryTree.readPointer)

	if readNode.value == nil {
		return nil, false
	}

	element := readNode.value
	readNode.value = nil
	binaryTree.readPointer = (binaryTree.readPointer + 1) % binaryTree.size

	fmt.Println("READ", element)
	binaryTree.Print()
	fmt.Println()

	return element, true
}

func (binaryTree *BinaryTree) Print() {
	binaryTree.print(binaryTree.root)
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
	}

	return nil
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

func (binaryTree *BinaryTree) print(node *binaryTreeNode) {
	if node == nil {
		return
	}

	readPointerNode := binaryTree.findNode(binaryTree.readPointer)
	writePointerNode := binaryTree.findNode(binaryTree.writePointer)
	node.print(readPointerNode, writePointerNode)

	binaryTree.print(node.leftNode)
	binaryTree.print(node.rightNode)
}
