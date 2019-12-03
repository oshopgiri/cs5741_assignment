package main

import (
	"github.com/oshopgiri/assignments/circular_buffer"
	"runtime"
	"sync"
)

func produceCircularBufferAsync(circularBufferOperations circular_buffer.ICircularBufferOperations, start, end int, waitGroup *sync.WaitGroup) {
	for i := start; i <= end; i++ {
		if ok := circularBufferOperations.Write(i); !ok {
			i--
		}
	}

	waitGroup.Done()
}

func consumeCircularBufferAsync(circularBufferOperations circular_buffer.ICircularBufferOperations, start, end int, waitGroup *sync.WaitGroup) {
	for i := start; i <= end; i++ {
		if _, ok := circularBufferOperations.Read(); !ok {
			i--
		}
	}

	waitGroup.Done()
}

func AsynchronousCircularBuffer(size int, count int, myCircularBuffer circular_buffer.CircularBuffer) {
	waitGroup := new(sync.WaitGroup)
	var circularBufferOperations circular_buffer.ICircularBufferOperations = &circular_buffer.CircularBufferOperations{}
	circularBufferOperations.Init(myCircularBuffer, size)

	threadsPerOperation := runtime.GOMAXPROCS(0) / 2
	inputsPerOperation := count / threadsPerOperation

	for i := 0; i < threadsPerOperation; i++ {
		start, end := i*inputsPerOperation+1, i*inputsPerOperation+inputsPerOperation+1
		if end > count {
			end = count
		}

		waitGroup.Add(1)
		go consumeCircularBufferAsync(circularBufferOperations, start, end, waitGroup)

		waitGroup.Add(1)
		go produceCircularBufferAsync(circularBufferOperations, start, end, waitGroup)
	}

	waitGroup.Wait()
}
