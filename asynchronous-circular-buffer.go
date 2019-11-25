package main

import (
	"github.com/oshopgiri/assignments/circular_buffer"
	"sync"
)

func produceCircularBufferAsync(myBuffer circular_buffer.CircularBuffer, count int, waitGroup *sync.WaitGroup) {
	stackWaitGroup := new(sync.WaitGroup)

	for i := 1; i <= count; i++ {
		stackWaitGroup.Add(1)
		go myBuffer.Write(i, stackWaitGroup)
	}

	stackWaitGroup.Wait()
	waitGroup.Done()
}

func consumeCircularBufferAsync(myBuffer circular_buffer.CircularBuffer, count int, waitGroup *sync.WaitGroup) {
	stackWaitGroup := new(sync.WaitGroup)

	for i := 1; i <= count; i++ {
		stackWaitGroup.Add(1)
		go myBuffer.Read(stackWaitGroup)
	}

	stackWaitGroup.Wait()
	waitGroup.Done()
}

func AsynchronousCircularBuffer(size int, count int, myBuffer circular_buffer.CircularBuffer) {
	waitGroup := new(sync.WaitGroup)
	myBuffer.Init(size)

	waitGroup.Add(1)
	go consumeCircularBufferAsync(myBuffer, count, waitGroup)

	waitGroup.Add(1)
	go produceCircularBufferAsync(myBuffer, count, waitGroup)

	waitGroup.Wait()
}
