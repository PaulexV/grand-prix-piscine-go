package piscine

func LastRune(s string) rune {

	var res rune = 0
	lenght := 0

	for range s {
		lenght++
	}
	res = []rune(s)[lenght-1]
	return res
}
