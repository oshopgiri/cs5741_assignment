package stack

import (
	"log"
	"sync"
)

type Array struct {
	sync.RWMutex
	elements []int
}

func (array *Array) InitArray() {
	array.elements = []int{}
}

func (array *Array) Push(element int) {
	array.Lock()
	defer array.Unlock()

	array.elements = append(array.elements, element)

	log.Println(array.elements, "<--", element)
}

func (array *Array) Pop() (int, bool) {
	if len(array.elements) > 0 {
		element := array.elements[len(array.elements)-1]

		array.Lock()

		array.elements = array.elements[:len(array.elements)-1]

		log.Println(array.elements, "-->", element)

		array.Unlock()

		return element, true
	} else {
		return 0, false
	}
}

func (array *Array) Print() {
	log.Println(array.elements)
}
