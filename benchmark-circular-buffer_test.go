package main

import (
	"github.com/oshopgiri/assignments/circular_buffer"
	"io/ioutil"
	"log"
	"testing"
)

func BenchmarkCircularBufferMain(b *testing.B) {
	log.SetOutput(ioutil.Discard)
}

func BenchmarkAsynchronousCircularBufferArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AsynchronousCircularBuffer(100, 10000, &circular_buffer.Array{})
	}
}

func BenchmarkSynchronousCircularBufferArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SynchronousCircularBuffer(100, 10000, &circular_buffer.Array{})
	}
}
