package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:] // skip program name at index [0]
	// condition 1 - no argument
	if len(args) == 0 {
		fmt.Println("File name missing")
		return
	}

	if len(args) > 1 { // condition 2 - more than 1 argument
		fmt.Println("Too many arguments")
		return
	}

	filename := args[0]
	content, err := os.ReadFile(filename)
	if err != nil {
		return
	}

	fmt.Print(string(content))
}
