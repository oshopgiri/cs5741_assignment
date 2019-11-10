package main

import (
	"github.com/oshopgiri/assignments/stack"
	"time"
)

func produceStackArray(myStack *stack.Array, n int) {
	for i := 1; i <= n; i++ {
		myStack.Push(i)
	}
}

func consumeStackArray(myStack *stack.Array, n int) {
	for i := 1; i <= n; i++ {
		if _, ok := myStack.Pop(); ok {
		} else {
			i--
			time.Sleep(1)
		}
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
