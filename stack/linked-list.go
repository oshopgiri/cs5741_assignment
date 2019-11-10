package stack

import (
	"fmt"
	"sync"
)

type LinkedListNode struct {
	value    int
	previous *LinkedListNode
	next     *LinkedListNode
}

type LinkedList struct {
	lastNode *LinkedListNode
	mutex    sync.RWMutex
}

func (linkedList *LinkedList) Push(element int) {
	linkedList.mutex.Lock()
	defer linkedList.mutex.Unlock()

	//fmt.Println("<--", element)

	node := &LinkedListNode{element, linkedList.lastNode, nil}
	if linkedList.lastNode != nil {
		linkedList.lastNode.next = node
	}
	linkedList.lastNode = node
}

func (linkedList *LinkedList) Pop() (int, bool) {
	linkedList.mutex.Lock()
	defer linkedList.mutex.Unlock()

	if linkedList.lastNode == nil {
		return 0, false
	} else {
		lastNode := linkedList.lastNode
		if lastNode.previous == nil {
			linkedList.lastNode = nil
		} else {
			linkedList.lastNode = lastNode.previous
			linkedList.lastNode.next = nil
		}

		//fmt.Println("-->", lastNode.value)

		return lastNode.value, true
	}
}

func (linkedList *LinkedList) Print() {
	listItems := ""
	currentNode := linkedList.lastNode
	for {
		if currentNode == nil {
			break
		}
		listItems = string(currentNode.value) + listItems
		currentNode = currentNode.previous
	}

	fmt.Println(listItems)
}
