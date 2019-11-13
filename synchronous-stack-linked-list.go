package main

import (
	"github.com/oshopgiri/assignments/stack"
)

func produceStackLinkedList(myStack *stack.LinkedList, n int) {
	for i := 1; i <= n; i++ {
		myStack.Push(i)
	}
}

func consumeStackLinkedList(myStack *stack.LinkedList, n int) {
	for i := 1; i <= n; i++ {
		myStack.Pop()
	}
}

func SynchronousStackLinkedList() {
	myStack := &stack.LinkedList{}
	n := 20

	produceStackLinkedList(myStack, n)
	consumeStackLinkedList(myStack, n)
}
