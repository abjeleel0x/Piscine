package piscine

func FirstWord(s string) string {
	word := ""
	for _, r := range s {
		if r == ' ' { // stop when first space found
			break
		}
		word += string(r)
	}
	return word + "\n"
}
