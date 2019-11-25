package stack

import (
	"log"
	"sync"
	"time"
)

type Array struct {
	elements []int
	sync.RWMutex
}

func (array *Array) Init() {
	array.elements = []int{}
}

func (array *Array) Push(element int, waitGroup *sync.WaitGroup) {
	if waitGroup != nil {
		defer waitGroup.Done()

		array.Lock()
		defer array.Unlock()
	}

	array.elements = append(array.elements, element)

	log.Println(array.elements, "<--", element)
}

func (array *Array) Pop(waitGroup *sync.WaitGroup) (int, bool) {
	if waitGroup != nil {
		defer waitGroup.Done()

		array.Lock()
		defer array.Unlock()
	}

	if len(array.elements) > 0 {
		element := array.elements[len(array.elements)-1]

		array.elements = array.elements[:len(array.elements)-1]

		log.Println(array.elements, "-->", element)

		return element, true
	} else {
		if waitGroup != nil {
			time.Sleep(1)

			waitGroup.Add(1)
			go array.Pop(waitGroup)
		}

		return 0, false
	}
}

func (array *Array) Print() {
	log.Println(array.elements)
}
