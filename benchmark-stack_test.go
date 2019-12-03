package main

import (
	"github.com/oshopgiri/assignments/stack"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func BenchmarkStackMain(b *testing.B) {
	log.SetOutput(ioutil.Discard)
	os.Stdout, _ = os.Open(os.DevNull)
}

func BenchmarkAsynchronousStackArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AsynchronousStack(100000, &stack.Array{})
	}
}

func BenchmarkSynchronousStackArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SynchronousStack(100000, &stack.Array{})
	}
}

func BenchmarkAsynchronousStackLinkedList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AsynchronousStack(100000, &stack.LinkedList{})
	}
}

func BenchmarkSynchronousStackLinkedList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SynchronousStack(100000, &stack.LinkedList{})
	}
}

func BenchmarkAsynchronousStackBinaryTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AsynchronousStack(100000, &stack.BinaryTree{})
	}
}

func BenchmarkSynchronousStackBinaryTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SynchronousStack(100000, &stack.BinaryTree{})
	}
}
