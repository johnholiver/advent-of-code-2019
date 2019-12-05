package computer

type IO struct {
	values []int
	cursor int
}

func NewIO(values []int) *IO {
	return &IO{
		values: values,
		cursor: 0,
	}
}

func (io *IO) Reset() {
	io.cursor = 0
}

func (io *IO) Next() {
	io.cursor++
}

func (io *IO) Read() int {
	value := io.values[io.cursor]
	io.Next()
	return value
}

func (io *IO) ReadAt(index int) int {
	return io.values[index]
}

func (io *IO) Write(value int) {
	io.values[io.cursor] = value
	io.Next()
}
