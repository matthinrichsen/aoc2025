package main

import "strconv"

type rotFunc func(int, int) int

func parseLine(line []byte) (rotFunc, int) {
	switch line[0] {
	case 'L':
		return left, getAmount(line[1:])
	case 'R':
		return right, getAmount(line[1:])
	default:
		panic(string(line))
	}
}

func getAmount(line []byte) int {
	res, err := strconv.Atoi(string(line))
	if err != nil {
		panic(err)
	}
	return res
}
