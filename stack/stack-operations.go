package stack

import (
	"fmt"
	"runtime"
)

type StackOperations struct {
	operations chan StackOperation
}

func (stackOperations *StackOperations) Init(stack Stack) {
	stackOperations.operations = make(chan StackOperation, runtime.GOMAXPROCS(0))
	stack.Init()

	go func() {
		for operation := range stackOperations.operations {
			operation(stack)
		}
	}()
}

func (stackOperations *StackOperations) Push(element interface{}) {
	printWaitChannel := make(chan bool)

	stackOperations.operations <- func(stack Stack) {
		stack.Push(element)

		fmt.Println("PUSH", element)
		stack.Print()
		fmt.Println()

		printWaitChannel <- true
	}

	<-printWaitChannel
}

func (stackOperations *StackOperations) Pop() (interface{}, bool) {
	responseChannel := make(chan interface{})
	statusChannel := make(chan bool)

	stackOperations.operations <- func(stack Stack) {
		element, ok := stack.Pop()

		if ok {
			fmt.Println("POP", element)
			stack.Print()
			fmt.Println()
		}

		responseChannel <- element
		statusChannel <- ok
	}

	return <-responseChannel, <-statusChannel
}

func (stackOperations *StackOperations) Close() {
	close(stackOperations.operations)
}
