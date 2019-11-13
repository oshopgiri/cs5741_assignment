package main

import (
	"github.com/oshopgiri/assignments/stack"
)

func produceStackArray(myStack *stack.Array, n int) {
	for i := 1; i <= n; i++ {
		myStack.Push(i)
	}
}

func consumeStackArray(myStack *stack.Array, n int) {
	for i := 1; i <= n; i++ {
		myStack.Pop()
	}
}

func synchronousStackArray() {
	myStack := &stack.Array{}
	myStack.InitArray()
	n := 20

	produceStackArray(myStack, n)
	consumeStackArray(myStack, n)
}

//func main() {
//	synchronousStackArray()
//}
