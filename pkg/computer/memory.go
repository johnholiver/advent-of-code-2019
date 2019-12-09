package computer

type Memory interface {
	Read(address int, mode ParamMode) int
	Write(address int, value int, mode ParamMode)
	String() string
}
