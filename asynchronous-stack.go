package main

import (
	"github.com/oshopgiri/assignments/stack"
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

func AsynchronousStack(count int, myStack stack.Stack, numberOfProducers, numberOfConsumers int) {
	waitGroup := new(sync.WaitGroup)
	var stackOperations stack.IStackOperations = &stack.StackOperations{}
	stackOperations.Init(myStack)

	consumerInputsPerOperation := count / numberOfConsumers
	if count > consumerInputsPerOperation*numberOfConsumers {
		consumerInputsPerOperation++
	}

	for i := 0; i < numberOfConsumers; i++ {
		start, end := i*consumerInputsPerOperation, i*consumerInputsPerOperation+consumerInputsPerOperation
		if end > count {
			end = count
		}

		waitGroup.Add(1)
		go consumeStackAsync(stackOperations, start, end, waitGroup)
	}

	producerInputsPerOperation := count / numberOfProducers
	if count > producerInputsPerOperation*numberOfProducers {
		producerInputsPerOperation++
	}

	for i := 0; i < numberOfProducers; i++ {
		start, end := i*producerInputsPerOperation, i*producerInputsPerOperation+producerInputsPerOperation
		if end > count {
			end = count
		}

		waitGroup.Add(1)
		go produceStackAsync(stackOperations, start, end, waitGroup)
	}

	waitGroup.Wait()
	stackOperations.Close()
}
