package main

import (
	"github.com/01-edu/z01"

	"os"
)

func main() {
	arg := os.Args
	lenS := 0
	for range arg {
		lenS++
	}
	for i := 1; i < lenS; i++ {
		for _, char := range arg[i] {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
