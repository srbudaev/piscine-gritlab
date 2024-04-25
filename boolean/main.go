package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isEven(nbr int) bool {
	return nbr%2 == 0
}

func main() {
	numArgs := len(os.Args) - 1

	if isEven(numArgs) {
		for _, r := range "I have an even number of arguments" {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	} else {
		for _, r := range "I have an odd number of arguments" {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
