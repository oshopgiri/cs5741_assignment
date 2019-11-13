package main

import (
	"github.com/oshopgiri/assignments/stack"
	"sync"
)

func produceStackLinkedListAsync(myStack *stack.LinkedList, n int, waitGroup *sync.WaitGroup) {
	for i := 1; i <= n; i++ {
		myStack.Push(i)
	}

	waitGroup.Done()
}

func consumeStackLinkedListAsync(myStack *stack.LinkedList, n int, waitGroup *sync.WaitGroup) {
	for i := 1; i <= n; i++ {
		if _, ok := myStack.Pop(); ok {
		} else {
			i--
		}
	}

	waitGroup.Done()
}

func AsynchronousStackLinkedList() {
	waitGroup := new(sync.WaitGroup)
	myStack := &stack.LinkedList{}
	n := 20

	waitGroup.Add(1)
	go consumeStackLinkedListAsync(myStack, n, waitGroup)

	waitGroup.Add(1)
	go produceStackLinkedListAsync(myStack, n, waitGroup)

	waitGroup.Wait()
}
