package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_part1(t *testing.T) {
	assert.Equal(t, true, isValid(111111))
	assert.Equal(t, false, isValid(223450))
	assert.Equal(t, false, isValid(123789))

}

func Test_part2(t *testing.T) {
	assert.Equal(t, true, isValid2(112233))
	assert.Equal(t, false, isValid2(123444))
	assert.Equal(t, true, isValid2(111122))
}
