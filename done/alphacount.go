package piscine

func AlphaCount(str string) int {

	res := 0

	for index := range str {

		if (str[index] >= 'A' && str[index] <= 'Z') || (str[index] >= 'a' && str[index] <= 'z') {

			res++
		}
		index++
	}
	return res
}
