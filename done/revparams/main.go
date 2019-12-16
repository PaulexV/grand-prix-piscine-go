package main

import (
	"github.com/01-edu/z01"

	"os"
)

func putStr(s string) {

	for _, char := range s {
		z01.PrintRune(char)
	}
}

func main() {

	arg := os.Args
	lenS := 0

	for range arg {
		lenS++
	}
	for i := lenS - 1; i > 0; i-- {
		putStr(arg[i] + "\n")
	}
}
