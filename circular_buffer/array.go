package circular_buffer

import (
	"fmt"
)

type Array struct {
	elements     []interface{}
	size         int
	readPointer  int
	writePointer int
}

func (array *Array) Init(size int) {
	array.elements = make([]interface{}, size)
	
	array.size = size
	array.readPointer = 0
	array.writePointer = 0
}

func (array *Array) Write(element interface{}) bool {
	if array.elements[array.writePointer] == nil {
		array.elements[array.writePointer] = element
		array.writePointer = array.nextIndex(array.writePointer)

		return true
	}

	return false
}

func (array *Array) Read() (interface{}, bool) {
	if array.elements[array.readPointer] != nil {
		element := array.elements[array.readPointer]
		array.elements[array.readPointer] = nil
		array.readPointer = array.nextIndex(array.readPointer)

		return element, true
	}

	return nil, false
}

func (array *Array) Print() {
	for i := 0; i < len(array.elements); i++ {
		fmt.Printf("%5v", array.elements[i])
		if i == array.readPointer {
			fmt.Print(" <-- readPointer")
		}

		if i == array.writePointer {
			fmt.Print(" <-- writePointer")
		}
		fmt.Println()
	}
}

func (array *Array) nextIndex(currentIndex int) int {
	return (currentIndex + 1) % array.size
}
