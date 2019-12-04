package main

import (
	"github.com/oshopgiri/assignments/avl_tree"
	"sync"
)

func produceAVLTreeAsync(avlTreeOperations avl_tree.IAVLTreeOperations, start, end int, waitGroup *sync.WaitGroup) {
	for i := start; i < end; i++ {
		avlTreeOperations.Insert(i)
	}

	waitGroup.Done()
}

func consumeAVLTreeAsync(avlTreeOperations avl_tree.IAVLTreeOperations, start, end int, waitGroup *sync.WaitGroup) {
	//for i := start; i < end; i++ {
	//	if _, ok := avlTreeOperations.Delete(i); !ok {
	//		i--
	//	}
	//}

	waitGroup.Done()
}

func AsynchronousAVLTree(count int, myAVLTree avl_tree.IAVLTree, numberOfProducers, numberOfConsumers int) {
	waitGroup := new(sync.WaitGroup)
	var avlTreeOperations avl_tree.IAVLTreeOperations = &avl_tree.AVLTreeOperations{}
	avlTreeOperations.Init(myAVLTree)

	consumerInputsPerOperation := count / numberOfConsumers
	if count > consumerInputsPerOperation*numberOfConsumers {
		consumerInputsPerOperation++
	}

	for i := 0; i < numberOfConsumers; i++ {
		start, end := i*consumerInputsPerOperation, i*consumerInputsPerOperation+consumerInputsPerOperation
		if end > count {
			end = count
		}

		waitGroup.Add(1)
		go consumeAVLTreeAsync(avlTreeOperations, start, end, waitGroup)
	}

	producerInputsPerOperation := count / numberOfProducers
	if count > producerInputsPerOperation*numberOfProducers {
		producerInputsPerOperation++
	}

	for i := 0; i < numberOfProducers; i++ {
		start, end := i*producerInputsPerOperation, i*producerInputsPerOperation+producerInputsPerOperation
		if end > count {
			end = count
		}

		waitGroup.Add(1)
		go produceAVLTreeAsync(avlTreeOperations, start, end, waitGroup)
	}

	waitGroup.Wait()
	avlTreeOperations.Close()
}
