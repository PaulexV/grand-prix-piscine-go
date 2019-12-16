package main

import (
	"github.com/01-edu/z01"

	"os"
)

func putStr(s string) {

	for _, c := range s {
		z01.PrintRune(c)
	}
}

func main() {

	args := os.Args
	sorted := true

	for sorted == true {
		sorted = false
		for i := range args {
			if i > 1 && args[i-1] > args[i] {
				sorted = true
				args[i-1], args[i] = args[i], args[i-1]
			}
		}
	}
	for _, c := range args[1:] {
		putStr(c + "\n")
	}
}
