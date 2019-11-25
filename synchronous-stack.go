package main

import (
	"github.com/oshopgiri/assignments/stack"
)

func produceStackArray(myStack stack.Stack, count int) {
	for i := 1; i <= count; i++ {
		myStack.Push(i, nil)
	}
}

func consumeStackArray(myStack stack.Stack, count int) {
	for i := 1; i <= count; i++ {
		myStack.Pop(nil)
	}
}

func SynchronousStack(count int, myStack stack.Stack) {
	myStack.Init()

	produceStackArray(myStack, count)
	consumeStackArray(myStack, count)
}
