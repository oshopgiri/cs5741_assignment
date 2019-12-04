package main

import (
	"github.com/oshopgiri/assignments/circular_buffer"
	"sync"
)

func produceCircularBufferAsync(circularBufferOperations circular_buffer.ICircularBufferOperations, start, end int, waitGroup *sync.WaitGroup) {
	for i := start; i < end; i++ {
		if ok := circularBufferOperations.Write(i); !ok {
			i--
		}
	}

	waitGroup.Done()
}

func consumeCircularBufferAsync(circularBufferOperations circular_buffer.ICircularBufferOperations, start, end int, waitGroup *sync.WaitGroup) {
	for i := start; i < end; i++ {
		if _, ok := circularBufferOperations.Read(); !ok {
			i--
		}
	}

	waitGroup.Done()
}

func AsynchronousCircularBuffer(size, count int, myCircularBuffer circular_buffer.CircularBuffer, numberOfProducers, numberOfConsumers int) {
	waitGroup := new(sync.WaitGroup)
	var circularBufferOperations circular_buffer.ICircularBufferOperations = &circular_buffer.CircularBufferOperations{}
	circularBufferOperations.Init(myCircularBuffer, size)

	consumerInputsPerOperation := count / numberOfConsumers
	if count > consumerInputsPerOperation*numberOfConsumers {
		consumerInputsPerOperation++
	}

	for i := 0; i < numberOfConsumers; i++ {
		start, end := i*consumerInputsPerOperation, i*consumerInputsPerOperation+consumerInputsPerOperation
		if end > count {
			end = count
		}

		waitGroup.Add(1)
		go consumeCircularBufferAsync(circularBufferOperations, start, end, waitGroup)
	}

	producerInputsPerOperation := count / numberOfProducers
	if count > producerInputsPerOperation*numberOfProducers {
		producerInputsPerOperation++
	}

	for i := 0; i < numberOfProducers; i++ {
		start, end := i*producerInputsPerOperation, i*producerInputsPerOperation+producerInputsPerOperation
		if end > count {
			end = count
		}

		waitGroup.Add(1)
		go produceCircularBufferAsync(circularBufferOperations, start, end, waitGroup)
	}

	waitGroup.Wait()
	circularBufferOperations.Close()
}
