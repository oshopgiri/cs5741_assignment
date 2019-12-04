package circular_buffer

import (
	"fmt"
	"runtime"
)

type CircularBufferOperations struct {
	operations chan CircularBufferOperation
}

func (circularBufferOperations *CircularBufferOperations) Init(circularBuffer CircularBuffer, size int) {
	circularBufferOperations.operations = make(chan CircularBufferOperation, runtime.GOMAXPROCS(0))
	circularBuffer.Init(size)

	go func() {
		for operation := range circularBufferOperations.operations {
			operation(circularBuffer)
		}
	}()
}

func (circularBufferOperations *CircularBufferOperations) Write(element interface{}) bool {
	statusChannel := make(chan bool)

	circularBufferOperations.operations <- func(circularBuffer CircularBuffer) {
		ok := circularBuffer.Write(element)

		if ok {
			fmt.Println("WRITE", element)
			circularBuffer.Print()
			fmt.Println()
		}

		statusChannel <- ok
	}

	return <-statusChannel
}

func (circularBufferOperations *CircularBufferOperations) Read() (interface{}, bool) {
	responseChannel := make(chan interface{})
	statusChannel := make(chan bool)

	circularBufferOperations.operations <- func(circularBuffer CircularBuffer) {
		element, ok := circularBuffer.Read()

		if ok {
			fmt.Println("READ", element)
			circularBuffer.Print()
			fmt.Println()
		}

		responseChannel <- element
		statusChannel <- ok
	}

	return <-responseChannel, <-statusChannel
}

func (circularBufferOperations *CircularBufferOperations) Close() {
	close(circularBufferOperations.operations)
}
