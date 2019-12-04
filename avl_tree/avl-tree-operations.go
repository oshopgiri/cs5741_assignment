package avl_tree

import (
	"fmt"
	"runtime"
)

type AVLTreeOperations struct {
	operations chan AVLTreeOperation
}

func (avlTreeOperations *AVLTreeOperations) Init(avlTree IAVLTree) {
	avlTreeOperations.operations = make(chan AVLTreeOperation, runtime.GOMAXPROCS(0))
	avlTree.Init()

	go func() {
		for operation := range avlTreeOperations.operations {
			operation(avlTree)
		}
	}()
}

func (avlTreeOperations *AVLTreeOperations) Insert(element int) {
	printWaitChannel := make(chan bool)

	avlTreeOperations.operations <- func(avlTree IAVLTree) {
		avlTree.Insert(element)

		fmt.Println("INSERT", element)
		avlTree.Print()
		fmt.Println()

		printWaitChannel <- true
	}

	<-printWaitChannel
}

func (avlTreeOperations *AVLTreeOperations) Delete(element int) (int, bool) {
	responseChannel := make(chan int)
	statusChannel := make(chan bool)

	avlTreeOperations.operations <- func(avlTree IAVLTree) {
		element, ok := avlTree.Delete(element)

		if ok {
			fmt.Println("DELETE", element)
			avlTree.Print()
			fmt.Println()
		}

		responseChannel <- element
		statusChannel <- ok
	}

	return <-responseChannel, <-statusChannel
}

func (avlTreeOperations *AVLTreeOperations) Close() {
	close(avlTreeOperations.operations)
}
