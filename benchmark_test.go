package main

import (
	"io/ioutil"
	"log"
	"testing"
)

func BenchmarkMain(b *testing.B) {
	log.SetOutput(ioutil.Discard)
}

func BenchmarkAsynchronousStackArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AsynchronousStackArray()
	}
}

func BenchmarkSynchronousStackArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SynchronousStackArray()
	}
}

func BenchmarkAsynchronousStackLinkedList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AsynchronousStackLinkedList()
	}
}

func BenchmarkSynchronousStackLinkedList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SynchronousStackLinkedList()
	}
}

func BenchmarkAsynchronousStackBinaryTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AsynchronousStackBinaryTree()
	}
}

func BenchmarkSynchronousStackBinaryTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SynchronousStackBinaryTree()
	}
}
