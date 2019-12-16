package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb() {

	var i rune
	var j rune
	var k rune

	i = '0'
	for i <= '7' {

		j = i + 1
		for j <= '8' {

			k = j + 1
			for k <= '9' {

				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(k)
				if i < '7' {

					z01.PrintRune(44)
					z01.PrintRune(32)
				} else {

					z01.PrintRune('\n')
				}
				k++
			}
			j++
		}
		i++
	}
}
