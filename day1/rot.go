package main

func left(num int, amount int) int {
	for amount > 100 {
		amount -= 100
	}
	if num < amount {
		return rot(num, 100-amount)
	}
	return rot(num, -amount)
}

func right(num int, amount int) int {
	for amount > 100 {
		amount -= 100
	}
	return rot(num, amount)
}

func rot(num, amount int) int {
	return (num + amount) % 100
}
