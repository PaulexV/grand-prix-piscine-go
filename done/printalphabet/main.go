package main

import (
	"github.com/01-edu/z01"
)

func main() {

	var alpha rune
	alpha = 97

	for alpha < 123 {

		z01.PrintRune(alpha)
		alpha++
	}
	z01.PrintRune('\n')
}
