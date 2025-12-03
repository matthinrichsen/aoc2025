package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	assert.Equal(t, result{pos: 82, revolutions: 1}, left(50, 68))
	assert.Equal(t, result{pos: 52, revolutions: 0}, left(82, 30))
	assert.Equal(t, result{pos: 0, revolutions: 0}, right(52, 48))
	assert.Equal(t, result{pos: 95, revolutions: 0}, left(0, 5))
	assert.Equal(t, result{pos: 55, revolutions: 1}, right(95, 60))
	assert.Equal(t, result{pos: 0, revolutions: 0}, left(55, 55))
	assert.Equal(t, result{pos: 99, revolutions: 0}, left(0, 1))
	assert.Equal(t, result{pos: 0, revolutions: 0}, left(99, 99))
	assert.Equal(t, result{pos: 14, revolutions: 0}, right(0, 14))
	assert.Equal(t, result{pos: 32, revolutions: 1}, left(14, 82))
}
