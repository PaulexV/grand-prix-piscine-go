package main

import (
	"github.com/01-edu/z01"
)

func main() {

	var nb rune
	nb = 48

	for nb < 58 {

		z01.PrintRune(nb)
		nb++
	}
	z01.PrintRune('\n')
}
