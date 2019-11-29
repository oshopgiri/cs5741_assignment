package stack

import "fmt"

type linkedListNode struct {
	value    interface{}
	previous *linkedListNode
	next     *linkedListNode
}

func (node *linkedListNode) print() string {
	return fmt.Sprintf("| %v |->", node.value)
}

type LinkedList struct {
	lastNode *linkedListNode
}

func (linkedList *LinkedList) Init() {}

func (linkedList *LinkedList) Push(element interface{}) {
	node := &linkedListNode{element, linkedList.lastNode, nil}
	if linkedList.lastNode != nil {
		linkedList.lastNode.next = node
	}
	linkedList.lastNode = node
}

func (linkedList *LinkedList) Pop() (interface{}, bool) {
	if linkedList.lastNode == nil {
		return nil, false
	}

	lastNode := linkedList.lastNode
	if lastNode.previous == nil {
		linkedList.lastNode = nil
	} else {
		linkedList.lastNode = lastNode.previous
		linkedList.lastNode.next = nil
	}

	return lastNode.value, true
}

func (linkedList *LinkedList) Print() {
	output := ""

	currentNode := linkedList.lastNode
	for {
		if currentNode == nil {
			break
		}

		output = currentNode.print() + output
		currentNode = currentNode.previous
	}

	fmt.Println(output)
}
