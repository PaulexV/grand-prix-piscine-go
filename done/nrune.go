package piscine

func NRune(s string, n int) rune {

	var res rune = 0
	lenght := 0

	for range s {
		lenght++
	}

	if n >= 1 && n <= lenght {
		res = []rune(s)[n-1]
	}
	return res
}
