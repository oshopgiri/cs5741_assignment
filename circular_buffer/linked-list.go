package circular_buffer

import (
	"fmt"
	"sync"
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
	sync.RWMutex
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

func (linkedList *LinkedList) Write(element interface{}, waitGroup *sync.WaitGroup) bool {
	if waitGroup != nil {
		defer waitGroup.Done()

		linkedList.Lock()
		defer linkedList.Unlock()
	}

	if linkedList.writePointer.value == nil {
		linkedList.writePointer.value = element
		linkedList.writePointer = linkedList.writePointer.next

		fmt.Println("WRITE", element)
		linkedList.Print()
		fmt.Println()

		return true
	} else {
		if waitGroup != nil {
			waitGroup.Add(1)
			go linkedList.Write(element, waitGroup)
		}

		return false
	}
}

func (linkedList *LinkedList) Read(waitGroup *sync.WaitGroup) (interface{}, bool) {
	if waitGroup != nil {
		defer waitGroup.Done()

		linkedList.Lock()
		defer linkedList.Unlock()
	}

	if linkedList.readPointer.value != nil {
		element := linkedList.readPointer.value
		linkedList.readPointer.value = nil
		linkedList.readPointer = linkedList.readPointer.next

		fmt.Println("READ", element)
		linkedList.Print()
		fmt.Println()

		return element, true
	} else {
		if waitGroup != nil {
			waitGroup.Add(1)
			go linkedList.Read(waitGroup)
		}

		return nil, false
	}
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
