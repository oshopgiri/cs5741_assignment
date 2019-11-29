package stack

import "fmt"

type StackOperations struct {
	operations chan StackOperation
}

func (stackOperations *StackOperations) Init(stack Stack) {
	stackOperations.operations = make(chan StackOperation)

	go func() {
		for operation := range stackOperations.operations {
			operation(stack)
		}
	}()
}

func (stackOperations *StackOperations) Push(element interface{}) {
	stackOperations.operations <- func(stack Stack) {
		stack.Push(element)
		fmt.Println("PUSH ", element)
		stack.Print()
		fmt.Println()
	}
}

func (stackOperations *StackOperations) Pop() bool {
	responseChannel := make(chan bool)

	stackOperations.operations <- func(stack Stack) {
		element, ok := stack.Pop()
		responseChannel <- ok

		if ok {
			fmt.Println("POP ", element)
			stack.Print()
			fmt.Println()
		}
	}

	return <-responseChannel
}

func (stackOperations *StackOperations) Close() {
	close(stackOperations.operations)
}
