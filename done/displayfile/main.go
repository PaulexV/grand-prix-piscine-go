package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	mylen := 0
	for range args {
		mylen++
	}

	if mylen < 1 {
		fmt.Println("File name missing")
	} else if mylen >= 2 {
		fmt.Println("Too many arguments")
	} else {
		data, err := ioutil.ReadFile(args[0])
		if err != nil {
			fmt.Println("File reading error", err)
			return
		}
		fmt.Println(string(data))
	}
}
