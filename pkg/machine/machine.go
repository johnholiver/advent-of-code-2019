package machine

type ProcessingStepFunc func(input *int) (output []int, done bool)

type Machine interface {
	Exec()
	ExecOneStep()
	SetDebugMode(d bool)
}

type AI interface {
	GetNextInput() *int
	LastOutput([]int)
	SetDebugMode(d bool)
}
