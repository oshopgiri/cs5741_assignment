package stack

type Stack interface {
	Init()
	Push(interface{})
	Pop() (interface{}, bool)
	Print()
}

type IStackOperations interface {
	Init(Stack)
	Push(interface{})
	Pop() bool
	Close()
}

type StackOperation func(Stack)
