package stack

import (
	"fmt"
	"sync"
)

type Array struct {
	elements []interface{}
	sync.RWMutex
}

func (array *Array) Init() {}

func (array *Array) Push(element interface{}, waitGroup *sync.WaitGroup) {
	if waitGroup != nil {
		defer waitGroup.Done()

		array.Lock()
		defer array.Unlock()
	}

	array.elements = append(array.elements, element)

	fmt.Println(array.elements, "<--", element)
}

func (array *Array) Pop(waitGroup *sync.WaitGroup) (interface{}, bool) {
	if waitGroup != nil {
		defer waitGroup.Done()

		array.Lock()
		defer array.Unlock()
	}

	if len(array.elements) > 0 {
		element := array.elements[len(array.elements)-1]

		array.elements = array.elements[:len(array.elements)-1]

		fmt.Println(array.elements, "-->", element)

		return element, true
	} else {
		if waitGroup != nil {
			waitGroup.Add(1)
			go array.Pop(waitGroup)
		}

		return nil, false
	}
}

func (array *Array) Print() {
	fmt.Println(array.elements)
}
