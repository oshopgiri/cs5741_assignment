package stack

import (
	"fmt"
)

type Array struct {
	elements []interface{}
}

func (array *Array) Init() {}

func (array *Array) Push(element interface{}) {
	array.elements = append(array.elements, element)
}

func (array *Array) Pop() (interface{}, bool) {
	if len(array.elements) > 0 {
		element := array.elements[len(array.elements)-1]
		array.elements = array.elements[:len(array.elements)-1]

		return element, true
	} else {
		return nil, false
	}
}

func (array *Array) Print() {
	fmt.Println(array.elements)
}
