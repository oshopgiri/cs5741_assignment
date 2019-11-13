package stack

import (
	"log"
	"strconv"
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
	node := &LinkedListNode{element, linkedList.lastNode, nil}

	linkedList.mutex.Lock()
	defer linkedList.mutex.Unlock()

	if linkedList.lastNode != nil {
		linkedList.lastNode.next = node
	}
	linkedList.lastNode = node

	log.Println("<--", element)
}

func (linkedList *LinkedList) Pop() (int, bool) {
	if linkedList.lastNode == nil {
		return 0, false
	} else {
		linkedList.mutex.Lock()

		lastNode := linkedList.lastNode
		if lastNode.previous == nil {
			linkedList.lastNode = nil
		} else {
			linkedList.lastNode = lastNode.previous
			linkedList.lastNode.next = nil
		}

		log.Println("-->", lastNode.value)

		linkedList.mutex.Unlock()

		return lastNode.value, true
	}
}

func (linkedList *LinkedList) Print() {
	listItems := ""
	currentNode := linkedList.lastNode
	for {
		if currentNode == nil {
			break
		} else {
			listItems = " --> " + listItems
		}
		listItems = strconv.Itoa(currentNode.value) + listItems
		currentNode = currentNode.previous
	}

	log.Println(listItems)
}

func (linkedListNode *LinkedListNode) Print() {

}
