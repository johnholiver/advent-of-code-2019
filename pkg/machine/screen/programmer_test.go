package screen

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var programmerInput = `A,B,C,B,A,C
R,8,R,8
R,4,R,4,R,8
L,6,L,2
y
`

func Test_programmableAI(t *testing.T) {
	ai := NewProgrammerAI(programmerInput)
	expected := []int{
		65, 44, 66, 44, 67, 44, 66, 44, 65, 44, 67, 10, //Main
		82, 44, 56, 44, 82, 44, 56, 10, //A
		82, 44, 52, 44, 82, 44, 52, 44, 82, 44, 56, 10, //B
		76, 44, 54, 44, 76, 44, 50, 10, //C
		121, 10, //feed
	}
	assert.Equal(t, expected, ai.(*Programmer).programToInput(programmerInput))
}
