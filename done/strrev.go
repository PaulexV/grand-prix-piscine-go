package piscine

func ft_strlen(str string) int {

	var res int
	for i, rune := range str {
		res++
		i = i
		rune = rune
	}
	return (res)

}

func StrRev(s string) string {
	runes := []rune(s)
	for i, j := 0, ft_strlen(s)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
