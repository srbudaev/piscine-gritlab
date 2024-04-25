package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	for i, w := range arg[0] {
		if i > 1 {
			z01.PrintRune(w)
		}
	}
	z01.PrintRune('\n')
}
