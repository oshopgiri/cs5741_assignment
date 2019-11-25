package stack

import "sync"

type Stack interface {
	Init()
	Push(interface{}, *sync.WaitGroup)
	Pop(*sync.WaitGroup) (interface{}, bool)
	Print()
}
