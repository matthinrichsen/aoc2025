package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input []byte

func main() {
	pos := 50
	var count int
	var zeros int

	buf := bytes.NewBuffer(input)
	scanner := bufio.NewScanner(buf)

	for scanner.Scan() {
		b := scanner.Bytes()
		fn, amt := parseLine(b)
		result := fn(pos, amt)
		pos = result.pos
		count += result.revolutions
		if pos == 0 {
			zeros++
		}
	}
	if pos == 0 {
		zeros++
	}
	fmt.Println("password:", zeros, count+zeros)
}
