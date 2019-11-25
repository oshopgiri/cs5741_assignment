package stack

import (
	"log"
	"sync"
)

type linkedListNode struct {
	value    interface{}
	previous *linkedListNode
	next     *linkedListNode
}

type LinkedList struct {
	lastNode *linkedListNode
	sync.RWMutex
}

func (linkedList *LinkedList) Init() {}

func (linkedList *LinkedList) Push(element interface{}, waitGroup *sync.WaitGroup) {
	if waitGroup != nil {
		defer waitGroup.Done()

		linkedList.Lock()
		defer linkedList.Unlock()
	}

	node := &linkedListNode{element, linkedList.lastNode, nil}

	if linkedList.lastNode != nil {
		linkedList.lastNode.next = node
	}

	linkedList.lastNode = node

	log.Println("<--", element)
}

func (linkedList *LinkedList) Pop(waitGroup *sync.WaitGroup) (interface{}, bool) {
	if waitGroup != nil {
		defer waitGroup.Done()

		linkedList.Lock()
		defer linkedList.Unlock()
	}

	if linkedList.lastNode == nil {
		if waitGroup != nil {
			waitGroup.Add(1)
			go linkedList.Pop(waitGroup)
		}

		return 0, false
	} else {
		lastNode := linkedList.lastNode
		if lastNode.previous == nil {
			linkedList.lastNode = nil
		} else {
			linkedList.lastNode = lastNode.previous
			linkedList.lastNode.next = nil
		}

		log.Println("-->", lastNode.value)

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
		listItems = currentNode.value.(string) + listItems
		currentNode = currentNode.previous
	}

	log.Println(listItems)
}

func (linkedListNode *linkedListNode) Print() {

}
