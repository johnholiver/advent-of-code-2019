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

type NoopAi struct{}

func NewNoopAI() AI {
	return &NoopAi{}
}
func (ai *NoopAi) GetNextInput() *int {
	return nil
}
func (ai *NoopAi) LastOutput(o []int) {
	return
}
func (ai *NoopAi) SetDebugMode(d bool) {
	return
}
