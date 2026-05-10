package main

import "fmt"

func FirstWord(s string) string {
	word := ""
	for _, r := range s {
		if r == ' ' {
			break
		}
		word += string(r)
	}
	return word + "\n"
}
