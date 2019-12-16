package piscine

func ToUpper(s string) string {

	str := []byte(s)

	for i := range str {

		if str[i] >= 'a' && str[i] <= 'z' {

			str[i] -= 'a' - 'A'
		}
	}
	return string(str)
}
