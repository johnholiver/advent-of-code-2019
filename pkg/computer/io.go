package computer

type IO interface {
	Reset()
	Previous()
	Next()
	CanRead() bool
	Read() int
	ReadAt(index int) int
	Set(values []int)
	Write(value int)
	WriteAt(value int, index int)
	Append(value int)
}
