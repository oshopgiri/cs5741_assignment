package circular_buffer

type CircularBuffer interface {
	Init(int)
	Write(interface{}) bool
	Read() (interface{}, bool)
	Print()
}

type ICircularBufferOperations interface {
	Init(CircularBuffer, int)
	Write(interface{}) bool
	Read() (interface{}, bool)
	Close()
}

type CircularBufferOperation func(CircularBuffer)
