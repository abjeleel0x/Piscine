package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// Generate the string for a specific quad type given dimensions x and y
func generateQuad(x, y int, typeQuad string) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var result strings.Builder

	for j := 1; j <= y; j++ {
		for i := 1; i <= x; i++ {
			if j == 1 { // Top row
				if i == 1 {
					result.WriteByte(getCorner(typeQuad, "topL"))
				} else if i == x {
					result.WriteByte(getCorner(typeQuad, "topR"))
				} else {
					result.WriteByte(getBorder(typeQuad, "h"))
				}
			} else if j == y { // Bottom row
				if i == 1 {
					result.WriteByte(getCorner(typeQuad, "botL"))
				} else if i == x {
					result.WriteByte(getCorner(typeQuad, "botR"))
				} else {
					result.WriteByte(getBorder(typeQuad, "h"))
				}
			} else { // Middle rows
				if i == 1 || i == x {
					result.WriteByte(getBorder(typeQuad, "v"))
				} else {
					result.WriteByte(' ')
				}
			}
		}
		result.WriteByte('\n')
	}
	return result.String()
}

func getCorner(q, pos string) byte {
	switch q {
	case "A":
		return 'o'
	case "B":
		if pos == "topL" || pos == "botR" {
			return '/'
		}
		return '\\'
	case "C":
		if pos == "topL" || pos == "topR" {
			return 'A'
		}
		return 'C'
	case "D":
		if pos == "topL" || pos == "botL" {
			return 'A'
		}
		return 'C'
	case "E":
		if pos == "topL" || pos == "botR" {
			return 'A'
		}
		return 'C'
	}
	return ' '
}

func getBorder(q, dir string) byte {
	if q == "A" || q == "B" {
		if dir == "h" {
			return '-'
		}
		return '|'
	}
	return 'B'
}

func main() {
	input, _ := io.ReadAll(os.Stdin)
	if len(input) == 0 {
		return
	}

	strInput := string(input)
	lines := strings.Split(strInput, "\n")
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	y := len(lines)
	if y == 0 {
		return
	}
	x := len(lines[0])

	var results []string
	quads := []string{"A", "B", "C", "D", "E"}

	for _, q := range quads {
		if generateQuad(x, y, q) == strInput {
			results = append(results, fmt.Sprintf("[quad%s] [%d] [%d]", q, x, y))
		}
	}

	if len(results) == 0 {
		fmt.Println("Not a quad function")
	} else {
		fmt.Println(strings.Join(results, " || "))
	}
}
