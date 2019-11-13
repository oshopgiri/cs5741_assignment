package circular_buffer

import (
	"fmt"
	"sync"
)

type Array struct {
	elements     []int
	size         int
	readPointer  int
	writePointer int
	mutex        sync.RWMutex
}

func (array *Array) InitArray(size int) {
	array.size = size
	array.elements = make([]int, size)
	for i := range array.elements {
		array.elements[i] = -1
	}
	array.readPointer = 0
	array.writePointer = 0
}

func (array *Array) Write(element int) bool {
	if (array.writePointer == array.readPointer) && (array.elements[array.writePointer] != -1) {
		return false
	} else {
		array.mutex.Lock()
		array.elements[array.writePointer] = element
		//fmt.Println(array.elements, "<--", element)
		array.writePointer = array.nextIndex(array.writePointer)
		array.mutex.Unlock()

		return true
	}
}

func (array *Array) Read() (int, bool) {
	if array.elements[array.readPointer] != 0 {
		element := array.elements[array.readPointer]

		array.mutex.Lock()
		array.elements[array.readPointer] = -1
		//fmt.Println(array.elements, "-->", element)
		array.readPointer = array.nextIndex(array.readPointer)
		array.mutex.Unlock()

		return element, true
	} else {
		return 0, false
	}
}

func (array *Array) Print() {
	fmt.Println(array.elements)
}

func (array *Array) nextIndex(currentIndex int) int {
	return (currentIndex + 1) % array.size
}
