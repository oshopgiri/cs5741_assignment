package main

import (
	"github.com/oshopgiri/assignments/circular_buffer"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func BenchmarkCircularBufferMain(b *testing.B) {
	log.SetOutput(ioutil.Discard)
	os.Stdout, _ = os.Open(os.DevNull)
}

func BenchmarkAsynchronousCircularBufferArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AsynchronousCircularBuffer(100, 100000, &circular_buffer.Array{})
	}
}

func BenchmarkSynchronousCircularBufferArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SynchronousCircularBuffer(100, 100000, &circular_buffer.Array{})
	}
}

func BenchmarkAsynchronousCircularBufferLinkedList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AsynchronousCircularBuffer(100, 100000, &circular_buffer.LinkedList{})
	}
}

func BenchmarkSynchronousCircularBufferLinkedList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SynchronousCircularBuffer(100, 100000, &circular_buffer.LinkedList{})
	}
}

func BenchmarkAsynchronousCircularBufferBinaryTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AsynchronousCircularBuffer(100, 100000, &circular_buffer.BinaryTree{})
	}
}

func BenchmarkSynchronousCircularBufferBinaryTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SynchronousCircularBuffer(100, 100000, &circular_buffer.BinaryTree{})
	}
}
