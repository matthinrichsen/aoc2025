package main

type result struct {
	pos         int
	revolutions int
}

func left(num int, amount int) result {
	passesZero := 0
	for amount >= 100 {
		if amount != 100 {
			passesZero++
		}
		amount -= 100
	}

	pos := num - amount
	if pos < 0 {
		if num != 0 && pos != 0 {
			passesZero++
		}
		pos += 100
	}
	return result{
		pos:         pos,
		revolutions: passesZero,
	}
}

func right(num int, amount int) result {
	passesZero := 0
	for amount >= 100 {
		if amount != 100 {
			passesZero++
		}
		amount -= 100
	}

	res := num + amount
	if res >= 100 {
		if res != 100 {
			passesZero++
		}
		res -= 100
	}
	return result{
		pos:         res,
		revolutions: passesZero,
	}
}
