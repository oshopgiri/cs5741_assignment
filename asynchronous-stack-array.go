package main

import (
	"github.com/oshopgiri/assignments/stack"
	"sync"
	"time"
)

func produceStackArrayAsync(myStack *stack.Array, n int, waitGroup *sync.WaitGroup) {
	for i := 1; i <= n; i++ {
		myStack.Push(i)
	}

	waitGroup.Done()
}

func consumeStackArrayAsync(myStack *stack.Array, n int, waitGroup *sync.WaitGroup) {
	for i := 1; i <= n; i++ {
		if _, ok := myStack.Pop(); ok {
		} else {
			i--
			time.Sleep(1)
		}
	}

	waitGroup.Done()
}

func asynchronousStackArray() {
	waitGroup := new(sync.WaitGroup)
	myStack := &stack.Array{}
	myStack.InitArray()
	n := 20

	waitGroup.Add(1)
	go consumeStackArrayAsync(myStack, n, waitGroup)

	waitGroup.Add(1)
	go produceStackArrayAsync(myStack, n, waitGroup)

	waitGroup.Wait()
}

//func main() {
//	asynchronousStackArray()
//}
