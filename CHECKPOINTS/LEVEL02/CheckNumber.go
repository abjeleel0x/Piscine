package piscine

func CheckNumber(arg string) bool {
	for _, r := range arg {       // loop through each character (rune)
		if r >= '0' && r <= '9' { // check if it's a digit
			return true
		}
	}
	return false // no digits found
}
