package circular_buffer

import (
	"fmt"
)

type linkedListNode struct {
	value    interface{}
	previous *linkedListNode
	next     *linkedListNode
}

func (node *linkedListNode) print() {
	fmt.Printf("| %5v |", node.value)
}

type LinkedList struct {
	readPointer  *linkedListNode
	writePointer *linkedListNode
}

func (linkedList *LinkedList) Init(size int) {
	firstNode := &linkedListNode{}
	lastNode := &linkedListNode{}

	for i := 0; i < size; i++ {
		node := &linkedListNode{nil, lastNode, nil}

		if i == 0 {
			firstNode = node
			lastNode = node
		} else {
			lastNode.next = node
			lastNode = node
		}
	}

	lastNode.next = firstNode
	firstNode.previous = lastNode

	linkedList.readPointer = firstNode
	linkedList.writePointer = firstNode
}

func (linkedList *LinkedList) Write(element interface{}) bool {
	if linkedList.writePointer.value == nil {
		linkedList.writePointer.value = element
		linkedList.writePointer = linkedList.writePointer.next

		return true
	}

	return false
}

func (linkedList *LinkedList) Read() (interface{}, bool) {
	if linkedList.readPointer.value != nil {
		element := linkedList.readPointer.value
		linkedList.readPointer.value = nil
		linkedList.readPointer = linkedList.readPointer.next

		return element, true
	}

	return nil, false
}

func (linkedList *LinkedList) Print() {
	currentNode := linkedList.readPointer
	startNode := currentNode

	for {
		currentNode.print()
		if currentNode == linkedList.readPointer {
			fmt.Print(" <-- readPointer")
		}
		if currentNode == linkedList.writePointer {
			fmt.Print(" <-- writePointer")
		}
		fmt.Println()

		if currentNode == startNode {
			break
		} else {
			fmt.Printf("%5v\n", "↓")
		}

		currentNode = currentNode.next
	}

	fmt.Println()
}
