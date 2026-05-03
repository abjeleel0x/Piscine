package main

import "github.com/01-edu/z01"

type point struct {
	x, y rune
}

func setPoint(ptr *point) {
	a := 'b' - 'a' // 1
	b := 'c' - 'a' // 2
	d := 'e' - 'a' // 4
	k := 'k' - 'a' // 10

	ptr.x = d*k + b // 4*10 + 2 = 42
	ptr.y = b*k + a // 2*10 + 1 = 21
}

func showRune(r rune) {
	z01.PrintRune(r)
}

func showStr(s string) {
	for _, ch := range s {
		showRune(ch)
	}
}

func showNum(n rune) {
	k := 'k' - 'a' // 10
	tens := n / k
	ones := n - tens*k
	zero := 'a' - 'a' + '0'
	showRune(tens + zero)
	showRune(ones + zero)
}

func main() {
	points := &point{}
	setPoint(points)

	showStr("x = ")
	showNum(points.x)
	showStr(", y = ")
	showNum(points.y)
	showStr("\n")
}
