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

func isEven() int {

	length := 0

	for i := range os.Args[1:] {
		length = i + 1
	}

	if length%2 == 0 {
		return 1
	}
	return 0
}

func main() {

	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"

	if isEven() == 1 {
		putStr(EvenMsg + "\n")
	} else {
		putStr(OddMsg + "\n")
	}
}
