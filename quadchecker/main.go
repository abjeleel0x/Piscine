package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func quadA(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var res strings.Builder
	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if row == 1 || row == y {
				if col == 1 || col == x {
					res.WriteRune('o')
				} else {
					res.WriteRune('-')
				}
			} else {
				if col == 1 || col == x {
					res.WriteRune('|')
				} else {
					res.WriteRune(' ')
				}
			}
		}
		res.WriteRune('\n')
	}
	return res.String()
}

func quadB(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var res strings.Builder
	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if row == 1 && col == 1 {
				res.WriteRune('/')
			} else if row == 1 && col == x {
				res.WriteRune('\\')
			} else if row == y && col == 1 {
				res.WriteRune('\\')
			} else if row == y && col == x {
				res.WriteRune('/')
			} else if row == 1 || row == y || col == 1 || col == x {
				res.WriteRune('*')
			} else {
				res.WriteRune(' ')
			}
		}
		res.WriteRune('\n')
	}
	return res.String()
}

func quadC(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var res strings.Builder
	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if row == 1 && (col == 1 || col == x) {
				res.WriteRune('A')
			} else if row == y && (col == 1 || col == x) {
				res.WriteRune('C')
			} else if row == 1 || row == y || col == 1 || col == x {
				res.WriteRune('B')
			} else {
				res.WriteRune(' ')
			}
		}
		res.WriteRune('\n')
	}
	return res.String()
}

func quadD(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var res strings.Builder
	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if (row == 1 && col == 1) || (row == y && col == 1) {
				res.WriteRune('A')
			} else if (row == 1 && col == x) || (row == y && col == x) {
				res.WriteRune('C')
			} else if row == 1 || row == y || col == 1 || col == x {
				res.WriteRune('B')
			} else {
				res.WriteRune(' ')
			}
		}
		res.WriteRune('\n')
	}
	return res.String()
}

func quadE(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var res strings.Builder
	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if row == 1 && (col == 1 || col == x) {
				res.WriteRune('A')
			} else if row == y && (col == 1 || col == x) {
				res.WriteRune('C')
			} else if row == 1 || row == y {
				res.WriteRune('B')
			} else if col == 1 || col == x {
				res.WriteRune('B')
			} else {
				res.WriteRune(' ')
			}
		}
		res.WriteRune('\n')
	}
	return res.String()
}

func main() {
	// Step 1: Read from stdin
	reader := bufio.NewReader(os.Stdin)
	input, _ := io.ReadAll(reader)
	content := string(input)

	if strings.TrimSpace(content) == "" {
		fmt.Println("Not a quad function")
		return
	}

	lines := strings.Split(strings.TrimRight(content, "\n"), "\n")
	y := len(lines)
	x := len([]rune(lines[0]))

	matches := []string{}

	// Step 2: Compare input to each quad
	if content == quadA(x, y) {
		matches = append(matches, fmt.Sprintf("[quadA] [%d] [%d]", x, y))
	}
	if content == quadB(x, y) {
		matches = append(matches, fmt.Sprintf("[quadB] [%d] [%d]", x, y))
	}
	if content == quadC(x, y) {
		matches = append(matches, fmt.Sprintf("[quadC] [%d] [%d]", x, y))
	}
	if content == quadD(x, y) {
		matches = append(matches, fmt.Sprintf("[quadD] [%d] [%d]", x, y))
	}
	if content == quadE(x, y) {
		matches = append(matches, fmt.Sprintf("[quadE] [%d] [%d]", x, y))
	}

	// Step 3: Print result
	if len(matches) == 0 {
		fmt.Println("Not a quad function")
	} else {
		fmt.Println(strings.Join(matches, " || "))
	}
}
