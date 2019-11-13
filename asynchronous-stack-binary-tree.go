package main

import (
	"github.com/oshopgiri/assignments/stack"
	"sync"
	"time"
)

func produceStackBinaryTreeAsync(myStack *stack.BinaryTree, n int, waitGroup *sync.WaitGroup) {
	for i := 1; i <= n; i++ {
		myStack.Push(i)
	}

	waitGroup.Done()
}

func consumeStackBinaryTreeAsync(myStack *stack.BinaryTree, n int, waitGroup *sync.WaitGroup) {
	for i := 1; i <= n; i++ {
		if _, ok := myStack.Pop(); ok {
		} else {
			i--
			time.Sleep(1)
		}
	}

	waitGroup.Done()
}

func AsynchronousStackBinaryTree() {
	waitGroup := new(sync.WaitGroup)
	myStack := &stack.BinaryTree{}
	n := 20

	waitGroup.Add(1)
	go consumeStackBinaryTreeAsync(myStack, n, waitGroup)

	waitGroup.Add(1)
	go produceStackBinaryTreeAsync(myStack, n, waitGroup)

	waitGroup.Wait()
}
