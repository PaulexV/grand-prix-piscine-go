package piscine

func ToLower(s string) string {

	str := []byte(s)

	for i := range str {

		if str[i] >= 'A' && str[i] <= 'Z' {

			str[i] += 'a' - 'A'
		}
	}
	return string(str)
}
