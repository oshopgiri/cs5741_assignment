package circular_buffer

import "sync"

type CircularBuffer interface {
	Init(int)
	Write(interface{}, *sync.WaitGroup) bool
	Read(*sync.WaitGroup) (interface{}, bool)
	Print()
}
