package main

import (
	"github.com/oshopgiri/assignments/stack"
)

func produceStackBinaryTree(myStack *stack.BinaryTree, n int) {
	for i := 1; i <= n; i++ {
		myStack.Push(i)
	}
}

func consumeStackBinaryTree(myStack *stack.BinaryTree, n int) {
	for i := 1; i <= n; i++ {
		myStack.Pop()
	}
}

func SynchronousStackBinaryTree() {
	myStack := &stack.BinaryTree{}
	n := 20

	produceStackBinaryTree(myStack, n)
	consumeStackBinaryTree(myStack, n)
}
