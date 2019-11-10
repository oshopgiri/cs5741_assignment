package main

import (
	"github.com/oshopgiri/assignments/stack"
	"time"
)

func produceStackLinkedList(myStack *stack.LinkedList, n int) {
	for i := 1; i <= n; i++ {
		myStack.Push(i)
	}
}

func consumeStackLinkedList(myStack *stack.LinkedList, n int) {
	for i := 1; i <= n; i++ {
		if _, ok := myStack.Pop(); ok {
		} else {
			i--
			time.Sleep(1)
		}
	}
}

func synchronousStackLinkedList() {
	myStack := &stack.LinkedList{}
	n := 20

	produceStackLinkedList(myStack, n)
	consumeStackLinkedList(myStack, n)
}

//func main() {
//	synchronousStackLinkedList()
//}
