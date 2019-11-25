package main

import (
	"github.com/oshopgiri/assignments/circular_buffer"
	"sync"
)

func produceCircularBuffer(myBuffer circular_buffer.CircularBuffer, count int, waitGroup *sync.WaitGroup) {
	for i := 1; i <= count; i++ {
		myBuffer.Write(i, nil)
	}
}

func consumeCircularBuffer(myBuffer circular_buffer.CircularBuffer, count int, waitGroup *sync.WaitGroup) {
	for i := 1; i <= count; i++ {
		myBuffer.Read(nil)
	}
}

func SynchronousCircularBuffer(size int, count int, myBuffer circular_buffer.CircularBuffer) {
	myBuffer.Init(size)

	produceCircularBuffer(myBuffer, count, nil)
	consumeCircularBuffer(myBuffer, count, nil)
}
