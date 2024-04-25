package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	setPoint(points)

	for _, r := range "x = " {
		z01.PrintRune(r)
	}

	printInteger(points.x)

	for _, r := range ", y = " {
		z01.PrintRune(r)
	}

	printInteger(points.y)

	z01.PrintRune('\n')
}

func printInteger(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	if n >= 10 {
		printInteger(n / 10)
	}
	z01.PrintRune(rune(n%10 + '0'))
}
