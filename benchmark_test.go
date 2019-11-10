package main

import "testing"

func BenchmarkAsynchronousStackArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		asynchronousStackArray()
	}
}

func BenchmarkSynchronousStackArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		synchronousStackArray()
	}
}

func BenchmarkAsynchronousStackLinkedList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		asynchronousStackLinkedList()
	}
}

func BenchmarkSynchronousStackLinkedList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		synchronousStackLinkedList()
	}
}

func BenchmarkAsynchronousStackBinaryTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		asynchronousStackBinaryTree()
	}
}

func BenchmarkSynchronousStackBinaryTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		synchronousStackBinaryTree()
	}
}
