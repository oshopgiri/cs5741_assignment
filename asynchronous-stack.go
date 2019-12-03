package main

import (
	"github.com/oshopgiri/assignments/stack"
	"runtime"
	"sync"
)

func produceStackAsync(stackOperations stack.IStackOperations, start, end int, waitGroup *sync.WaitGroup) {
	for i := start; i < end; i++ {
		stackOperations.Push(i)
	}

	waitGroup.Done()
}

func consumeStackAsync(stackOperations stack.IStackOperations, start, end int, waitGroup *sync.WaitGroup) {
	for i := start; i < end; i++ {
		if _, ok := stackOperations.Pop(); !ok {
			i--
		}
	}

	waitGroup.Done()
}

func AsynchronousStack(count int, myStack stack.Stack) {
	waitGroup := new(sync.WaitGroup)
	var stackOperations stack.IStackOperations = &stack.StackOperations{}
	stackOperations.Init(myStack)

	threadsPerOperation := runtime.GOMAXPROCS(0) / 2
	inputsPerOperation := count / threadsPerOperation
	if count > inputsPerOperation*threadsPerOperation {
		inputsPerOperation++
	}

	for i := 0; i < threadsPerOperation; i++ {
		start, end := i*inputsPerOperation, i*inputsPerOperation+inputsPerOperation
		if end > count {
			end = count
		}

		waitGroup.Add(1)
		go consumeStackAsync(stackOperations, start, end, waitGroup)

		waitGroup.Add(1)
		go produceStackAsync(stackOperations, start, end, waitGroup)

	}

	waitGroup.Wait()
}
