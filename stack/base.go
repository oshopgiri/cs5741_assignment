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
	Pop() (interface{}, bool)
	Close()
}

type StackOperation func(Stack)
