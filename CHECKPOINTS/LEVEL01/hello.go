package main

import "github.com/01-edu/z01"

func main() {
	str := "Hello World!" // save the string in a variable of your choice
	for _, char := range str { // loop through the string and
		z01.PrintRune(char) // print each rune character 1 by 1
	}
	z01.PrintRune('\n')
}
