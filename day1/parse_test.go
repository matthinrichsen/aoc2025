package main

import (
	"bufio"
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	inputBytes, err := os.ReadFile(`input.txt`)
	require.NoError(t, err)

	buf := bytes.NewBuffer(inputBytes)

	scanner := bufio.NewScanner(buf)

	for scanner.Scan() {
		assert.NotPanics(t, func() {
			parseLine(scanner.Bytes())
		})
	}
}
