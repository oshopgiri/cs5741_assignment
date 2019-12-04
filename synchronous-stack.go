package main

import (
	"github.com/oshopgiri/assignments/stack"
)

func produceStackArray(stackOperations stack.IStackOperations, count int) {
	for i := 1; i <= count; i++ {
		stackOperations.Push(i)
	}
}

func consumeStackArray(stackOperations stack.IStackOperations, count int) {
	for i := 1; i <= count; i++ {
		stackOperations.Pop()
	}
}

func SynchronousStack(count int, myStack stack.Stack) {
	var stackOperations stack.IStackOperations = &stack.StackOperations{}
	stackOperations.Init(myStack)

	produceStackArray(stackOperations, count)
	consumeStackArray(stackOperations, count)

	stackOperations.Close()
}
