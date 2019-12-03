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
	firstNode    *linkedListNode
	lastNode     *linkedListNode
	readPointer  *linkedListNode
	writePointer *linkedListNode
}

func (linkedList *LinkedList) Init(size int) {
	for i := 0; i < size; i++ {
		node := &linkedListNode{nil, linkedList.lastNode, nil}

		if linkedList.firstNode == nil {
			linkedList.firstNode = node
			linkedList.lastNode = node
		} else {
			linkedList.lastNode.next = node
			linkedList.lastNode = node
		}
	}

	linkedList.lastNode.next = linkedList.firstNode
	linkedList.firstNode.previous = linkedList.lastNode

	linkedList.readPointer = linkedList.firstNode
	linkedList.writePointer = linkedList.firstNode
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
	currentNode := linkedList.firstNode

	for {
		currentNode.print()
		if currentNode == linkedList.readPointer {
			fmt.Print(" <-- readPointer")
		}
		if currentNode == linkedList.writePointer {
			fmt.Print(" <-- writePointer")
		}
		fmt.Println()

		if currentNode == linkedList.lastNode {
			break
		} else {
			fmt.Printf("%5v\n", "↓")
		}

		currentNode = currentNode.next
	}

	fmt.Println()
}
