package main

import (
	"github.com/oshopgiri/assignments/stack"
	"runtime"
	"sync"
)

func produceStackAsync(stackOperations stack.IStackOperations, start, end int, waitGroup *sync.WaitGroup) {
	for i := start; i <= end; i++ {
		stackOperations.Push(i)
	}

	waitGroup.Done()
}

func consumeStackAsync(stackOperations stack.IStackOperations, start, end int, waitGroup *sync.WaitGroup) {
	for i := start; i <= end; i++ {
		if ok := stackOperations.Pop(); !ok {
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

	for i := 0; i < threadsPerOperation; i++ {
		waitGroup.Add(1)
		go consumeStackAsync(stackOperations, i*inputsPerOperation+1, i*inputsPerOperation+inputsPerOperation+1, waitGroup)

		waitGroup.Add(1)
		go produceStackAsync(stackOperations, i*inputsPerOperation+1, i*inputsPerOperation+inputsPerOperation+1, waitGroup)

	}

	waitGroup.Wait()
}
