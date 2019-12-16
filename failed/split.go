package main

import "fmt"

func contains(str, charset string, index, length int) bool {

	if length == 0 {
		return false
	}
	for i := 0; i < length; i++ {
		if str[index+i] == 0 || str[index+i] != charset[i] {
			return false
		}
	}
	return true
}

func myLen(s string) int {

	var length int

	for range s {
		length++
	}
	return length
}

func Split(str, charset string) []string {

	var (
		size   int
		resI   int
		lenStr int
	)
	length := myLen(charset)

	for i := range str {
		lenStr++
		if contains(str, charset, i, length) {
			size++
		}
	}
	if size == 0 {
		size++
	}
	res := make([]string, size)
	for i := 0; i < lenStr; i++ {
		if i-length >= 0 && contains(str, charset, i, length) && contains(str, charset, i-length, length) {
			i += length - 1
		} else if !contains(str, charset, i, length) {
			res[resI] += string(str[i])
		} else if contains(str, charset, i, length) && myLen(res[resI]) == 0 {
			i += length - 1
		} else {
			i += length - 1
			resI++
		}
	}
	return res
}

func main() {
	str := "salut"
	fmt.Println(Split(str, ""))
}
