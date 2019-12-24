package piscine

func getLen(arr []int) int {

	length := 0

	for range arr {
		length++
	}
	return length
}

func Map(f func(int) bool, arr []int) []bool {

	var tab []bool

	tab = make([]bool, getLen(arr))

	for k, v := range arr {
		tab[k] = f(v)
	}
	return tab
}
