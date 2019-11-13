package circular_buffer

type CircularBuffer interface {
	Write(int) bool
	Read() (int, bool)
	Print()
}
