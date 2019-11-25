package stack

import "sync"

type Stack interface {
	Init()
	Push(int, *sync.WaitGroup)
	Pop(*sync.WaitGroup) (int, bool)
	Print()
}
