package circular_buffer

import (
	"fmt"
	"log"
	"sync"
)

type Array struct {
	elements     []interface{}
	size         int
	readPointer  int
	writePointer int
	sync.RWMutex
}

func (array *Array) Init(size int) {
	array.size = size
	array.elements = make([]interface{}, size)
	//for i := range array.elements {
	//	array.elements[i] = nil
	//}
	array.readPointer = 0
	array.writePointer = 0
}

func (array *Array) Write(element interface{}, waitGroup *sync.WaitGroup) bool {
	if waitGroup != nil {
		defer waitGroup.Done()

		array.Lock()
		defer array.Unlock()
	}

	if (array.elements[array.writePointer] == nil) {
		array.elements[array.writePointer] = element
		array.writePointer = array.nextIndex(array.writePointer)

		log.Println(array.elements, "<--", element)

		return true
	} else {
		if waitGroup != nil {
			waitGroup.Add(1)
			go array.Write(element, waitGroup)
		}

		return false
	}
}

func (array *Array) Read(waitGroup *sync.WaitGroup) (interface{}, bool) {
	if waitGroup != nil {
		defer waitGroup.Done()

		array.Lock()
		defer array.Unlock()
	}

	if array.elements[array.readPointer] != nil {
		element := array.elements[array.readPointer]
		array.elements[array.readPointer] = nil
		array.readPointer = array.nextIndex(array.readPointer)

		log.Println(array.elements, "-->", element)

		return element, true
	} else {
		if waitGroup != nil {
			waitGroup.Add(1)
			go array.Read(waitGroup)
		}

		return 0, false
	}
}

func (array *Array) Print() {
	fmt.Println(array.elements)
}

func (array *Array) nextIndex(currentIndex int) int {
	return (currentIndex + 1) % array.size
}
