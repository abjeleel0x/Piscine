package main

import (
	"bufio"
	"io"
	"os"

	"github.com/01-edu/z01"
)

// printString prints a string rune by rune using z01.PrintRune
func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func main() {
	args := os.Args[1:]

	// No arguments: read from stdin
	if len(args) == 0 {
		reader := bufio.NewReader(os.Stdin)
		for {
			r, _, err := reader.ReadRune()
			if err == io.EOF {
				break
			}
			if err != nil {
				printString("ERROR: " + err.Error() + "\n")
				os.Exit(1)
			}
			z01.PrintRune(r)
		}
		return
	}

	// If file(s) provided
	for _, fileName := range args {
		file, err := os.Open(fileName)
		if err != nil {
			printString("ERROR: " + err.Error() + "\n")
			os.Exit(1)
		}

		reader := bufio.NewReader(file)
		for {
			r, _, err := reader.ReadRune()
			if err == io.EOF {
				break
			}
			if err != nil {
				printString("ERROR: " + err.Error() + "\n")
				file.Close()
				os.Exit(1)
			}
			z01.PrintRune(r)
		}
		file.Close()
	}
}
