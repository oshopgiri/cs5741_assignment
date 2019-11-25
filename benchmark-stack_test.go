package main

import (
	"github.com/oshopgiri/assignments/stack"
	"io/ioutil"
	"log"
	"testing"
)

func BenchmarkStackMain(b *testing.B) {
	log.SetOutput(ioutil.Discard)
}

func BenchmarkAsynchronousStackArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AsynchronousStack(10000, &stack.Array{})
	}
}

func BenchmarkSynchronousStackArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SynchronousStack(10000, &stack.Array{})
	}
}

func BenchmarkAsynchronousStackLinkedList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AsynchronousStack(10000, &stack.LinkedList{})
	}
}

func BenchmarkSynchronousStackLinkedList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SynchronousStack(10000, &stack.LinkedList{})
	}
}

func BenchmarkAsynchronousStackBinaryTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AsynchronousStack(10000, &stack.BinaryTree{})
	}
}

func BenchmarkSynchronousStackBinaryTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SynchronousStack(10000, &stack.BinaryTree{})
	}
}
