package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	// 1. Handle the sign

	if n < 0 {
		z01.PrintRune('-')
	} else if n == 0 {

		z01.PrintRune('0')

		return

	} else {
		// We convert positive to negative to handle the

		// full range of int (Negative range is larger than positive)

		n = -n
	}

	// 2. Recursive call to print preceding digits

	printRecursive(n)
}

func printRecursive(n int) {
	if n <= -10 {
		printRecursive(n / 10)
	}

	// 3. Convert digit to positive rune

	// Since n is negative, n%10 is negative or zero.

	// We subtract it from '0' to get the correct character.

	digit := '0' - rune(n%10)

	z01.PrintRune(digit)
}
