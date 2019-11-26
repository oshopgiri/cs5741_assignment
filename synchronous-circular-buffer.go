package main

import (
	"github.com/oshopgiri/assignments/circular_buffer"
	"sync"
)

func produceCircularBuffer(myBuffer circular_buffer.CircularBuffer, start int, end int, waitGroup *sync.WaitGroup) {
	for i := start; i <= end; i++ {
		myBuffer.Write(i, nil)
	}
}

func consumeCircularBuffer(myBuffer circular_buffer.CircularBuffer, start int, end int, waitGroup *sync.WaitGroup) {
	for i := start; i <= end; i++ {
		myBuffer.Read(nil)
	}
}

func SynchronousCircularBuffer(size int, count int, myBuffer circular_buffer.CircularBuffer) {
	myBuffer.Init(size)

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

		produceCircularBuffer(myBuffer, start, end, nil)
		consumeCircularBuffer(myBuffer, start, end, nil)
	}
}
