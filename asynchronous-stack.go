package main

import (
	"github.com/oshopgiri/assignments/stack"
	"sync"
)

func produceStackArrayAsync(myStack stack.Stack, count int, waitGroup *sync.WaitGroup) {
	stackWaitGroup := new(sync.WaitGroup)

	for i := 1; i <= count; i++ {
		stackWaitGroup.Add(1)
		go myStack.Push(i, stackWaitGroup)
	}

	stackWaitGroup.Wait()
	waitGroup.Done()
}

func consumeStackArrayAsync(myStack stack.Stack, count int, waitGroup *sync.WaitGroup) {
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
	go consumeStackArrayAsync(myStack, count, waitGroup)

	waitGroup.Add(1)
	go produceStackArrayAsync(myStack, count, waitGroup)

	waitGroup.Wait()
}
