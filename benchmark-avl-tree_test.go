package main

import (
	"github.com/oshopgiri/assignments/avl_tree"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func BenchmarkAVLTreeMain(b *testing.B) {
	log.SetOutput(ioutil.Discard)
	os.Stdout, _ = os.Open(os.DevNull)
}

func BenchmarkAsynchronousAVLTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AsynchronousAVLTree(100000, &avl_tree.AVLTree{})
	}
}
