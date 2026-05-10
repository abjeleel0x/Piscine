package piscine

import "github.com/01-edu/z01"

func Rot14(s string) string {
	result := ""
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			result += string('a' + (ch-'a'+14)%26)
		} else if ch >= 'A' && ch <= 'Z' {
			result += string('A' + (ch-'A'+14)%26)
		} else {
			result += string(ch)
		}
	}
	return result
}

func main() {
	result := Rot14("Hello! How are You?")

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
