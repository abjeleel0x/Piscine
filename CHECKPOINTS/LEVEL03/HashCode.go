Solution-1

package piscine

func HashCode(str string) string {
	result := ""
	for i := 0; i < len(str); i++ {
		char := str[i]
		hashValue := (int(char) + len(str)) % 127
		if hashValue < 33 {
			hashValue += 33
		}
		result += string(hashValue)
	}
	return result
}


Solutin-2

package piscine

func HashCode(dec string) string {
	if len(dec) == 0 {
		return ""
	}

	size := len(dec)
	var result string

	for _, r := range dec {
		hashed := (int(r) + size) % 127
		if hashed < 33 {
			hashed += 33
		}
		result += string(rune(hashed))
	}

	return result
}
