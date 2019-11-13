package stack

import (
	"fmt"
	"sync"
)

type Array struct {
	elements []int
	mutex    sync.RWMutex
}

func (array *Array) InitArray() {
	array.elements = []int{}
}

func (array *Array) Push(element int) {
	array.mutex.Lock()
	defer array.mutex.Unlock()

	//fmt.Println(array.elements, "<--", element)

	array.elements = append(array.elements, element)
}

func (array *Array) Pop() (int, bool) {
	if len(array.elements) > 0 {
		element := array.elements[len(array.elements)-1]

		array.mutex.Lock()
		array.elements = array.elements[:len(array.elements)-1]
		//fmt.Println(array.elements, "-->", element)
		array.mutex.Unlock()

		return element, true
	} else {
		return 0, false
	}
}

func (array *Array) Print() {
	fmt.Println(array.elements)
}
