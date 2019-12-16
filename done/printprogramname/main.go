package main

import (
	"github.com/01-edu/z01"

	"os"
)

func main() {

	arg := os.Args[0]

	for _, letters := range arg {

		z01.PrintRune(letters)
	}
	z01.PrintRune('\n')
}
