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

	buf := bytes.NewBuffer(input)
	scanner := bufio.NewScanner(buf)

	for scanner.Scan() {
		b := scanner.Bytes()
		fn, amt := parseLine(b)
		pos = fn(pos, amt)
		if pos == 0 {
			count++
		}
	}
	fmt.Println("password:", count)
}
