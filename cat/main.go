package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 1 {
		buffer := make([]byte, 1)
		for {
			_, err := os.Stdin.Read(buffer)
			if err != nil {
				break
			}
			z01.PrintRune(rune(buffer[0]))
		}
	} else {
		for _, filename := range os.Args[1:] {
			file, err := os.Open(filename)
			if err != nil {
				printError(err)
				os.Exit(1)
			}
			defer file.Close()

			buffer := make([]byte, 1)
			for {
				_, err := file.Read(buffer)
				if err != nil {
					break
				}
				z01.PrintRune(rune(buffer[0]))
			}
		}
	}
}

func printError(err error) {
	for _, r := range "ERROR: " {
		z01.PrintRune(r)
	}
	for _, r := range err.Error() {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
