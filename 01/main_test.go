package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_recursiveFuelPerModule(t *testing.T) {
	assert.Equal(t, 2, recursiveFuelPerModule(14))
	assert.Equal(t, 966, recursiveFuelPerModule(1969))
	assert.Equal(t, 50346, recursiveFuelPerModule(100756))
}
