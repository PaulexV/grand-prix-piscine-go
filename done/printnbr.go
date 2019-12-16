package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		if n/10 != 0 {
			PrintNbr(n / -10)
		}
		res := '0'
		for i := 0; i < -(n % 10); i++ {
			res++
		}
		z01.PrintRune(res)
	} else if n == 0 {
		z01.PrintRune('0')
	} else {
		if n/10 != 0 {
			PrintNbr(n / 10)
		}
		res := '0'
		for i := 0; i < n%10; i++ {
			res++
		}
		z01.PrintRune(res)
	}
}
