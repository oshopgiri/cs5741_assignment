package main

import (
	"github.com/oshopgiri/assignments/stack"
	"sync"
)

func produceStackAsync(myStack stack.Stack, count int, waitGroup *sync.WaitGroup) {
	stackWaitGroup := new(sync.WaitGroup)

	for i := 1; i <= count; i++ {
		stackWaitGroup.Add(1)
		go myStack.Push(i, stackWaitGroup)
	}

	stackWaitGroup.Wait()
	waitGroup.Done()
}

func consumeStackAsync(myStack stack.Stack, count int, waitGroup *sync.WaitGroup) {
	stackWaitGroup := new(sync.WaitGroup)

	for i := 1; i <= count; i++ {
		stackWaitGroup.Add(1)
		go myStack.Pop(stackWaitGroup)
	}

	stackWaitGroup.Wait()
	waitGroup.Done()
}

func AsynchronousStack(count int, myStack stack.Stack) {
	waitGroup := new(sync.WaitGroup)
	myStack.Init()

	waitGroup.Add(1)
	go consumeStackAsync(myStack, count, waitGroup)

	waitGroup.Add(1)
	go produceStackAsync(myStack, count, waitGroup)

	waitGroup.Wait()
}
