package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeft(t *testing.T) {
	assert.Equal(t, 82, left(50, 68))
	assert.Equal(t, 52, left(82, 30))
	assert.Equal(t, 0, right(52, 48))
	assert.Equal(t, 95, left(0, 5))
	assert.Equal(t, 55, right(95, 60))
	assert.Equal(t, 0, left(55, 55))
	assert.Equal(t, 99, left(0, 1))
	assert.Equal(t, 0, left(99, 99))
	assert.Equal(t, 14, right(0, 14))
	assert.Equal(t, 32, left(14, 82))
}
