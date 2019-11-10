package main

import (
	"github.com/oshopgiri/assignments/stack"
	"time"
)

func produceStackBinaryTree(myStack *stack.BinaryTree, n int) {
	for i := 1; i <= n; i++ {
		myStack.Push(i)
	}
}

func consumeStackBinaryTree(myStack *stack.BinaryTree, n int) {
	for i := 1; i <= n; i++ {
		if _, ok := myStack.Pop(); ok {
		} else {
			i--
			time.Sleep(1)
		}
	}
}

func synchronousStackBinaryTree() {
	myStack := &stack.BinaryTree{}
	n := 20

	produceStackBinaryTree(myStack, n)
	consumeStackBinaryTree(myStack, n)
}

//func main() {
//	synchronousStackBinaryTree()
//}