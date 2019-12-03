package main

import (
	"github.com/oshopgiri/assignments/circular_buffer"
)

func produceCircularBuffer(circularBufferOperations circular_buffer.ICircularBufferOperations, start int, end int) {
	for i := start; i <= end; i++ {
		circularBufferOperations.Write(i)
	}
}

func consumeCircularBuffer(circularBufferOperations circular_buffer.ICircularBufferOperations, start int, end int) {
	for i := start; i <= end; i++ {
		circularBufferOperations.Read()
	}
}

func SynchronousCircularBuffer(size int, count int, myCircularBuffer circular_buffer.CircularBuffer) {
	var circularBufferOperations circular_buffer.ICircularBufferOperations = &circular_buffer.CircularBufferOperations{}
	circularBufferOperations.Init(myCircularBuffer, size)

	loopCount := count / size
	if count%size > 0 {
		loopCount++
	}

	for i := 0; i < loopCount; i++ {
		start := i*size + 1
		end := i*size + size
		if end > count {
			end = count
		}

		produceCircularBuffer(circularBufferOperations, start, end)
		consumeCircularBuffer(circularBufferOperations, start, end)
	}
}
