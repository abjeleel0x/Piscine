package main

import (
	"fmt"
	"os"
)

func atoi(s string) int {
	n := 0

	for _, c := range s {

		if c < '0' || c > '9' {
			return -1
		}

		n = n*10 + int(c-'0')

	}

	return n
}

func main() {
	args := os.Args[1:]

	if len(args) < 3 || args[0] != "-c" {
		os.Exit(1)
	}

	count := atoi(args[1])

	if count < 0 {
		os.Exit(1)
	}

	files := args[2:]

	multiple := len(files) > 1

	hasError := false

	for i, file := range files {

		data, err := os.ReadFile(file)
		if err != nil {

			fmt.Println(err)

			hasError = true

			continue

		}

		if multiple {

			if i > 0 {
				fmt.Println()
			}

			fmt.Printf("==> %s <==\n", file)

		}

		start := len(data) - count

		if start < 0 {
			start = 0
		}

		fmt.Print(string(data[start:]))

	}

	if hasError {
		os.Exit(1)
	}
}
