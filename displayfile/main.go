package main

import (
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]

	switch len(args) {
	case 0:
		os.Stdout.Write([]byte("File name missing\n"))
	case 1:
		content, err := ioutil.ReadFile(args[0])
		if err != nil {
			os.Stdout.Write([]byte("Error reading file\n"))
			return
		}
		os.Stdout.Write(content)
	default:
		os.Stdout.Write([]byte("Too many arguments\n"))
	}
}
