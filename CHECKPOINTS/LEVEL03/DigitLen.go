package piscine

func DigitLen(n, base int) int {
	if base < 2 || base > 36 { // base must be valid
		return -1
	}
	if n < 0 { // if negative, make it positive
		n = -n
	}
	count := 0
	for {
		count++
		n = n / base
		if n == 0 { // stop when it reaches zero
			break
		}
	}
	return count
}
