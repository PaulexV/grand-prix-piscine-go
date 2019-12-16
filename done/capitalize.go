package piscine

func isAlnum(b byte) bool {

	if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
		return true
	}

	if b >= '0' && b <= '9' {
		return true
	}

	return false
}

func Capitalize(s string) string {

	str := []byte(s)

	for i := range str {
		if str[i] >= 'A' && str[i] <= 'Z' {
			str[i] += 'a' - 'A'
		}
		if i > 0 && !isAlnum(str[i-1]) && str[i] >= 'a' && str[i] <= 'z' {
			str[i] -= 'a' - 'A'
		}
	}
	if str[0] >= 'a' && str[0] <= 'z' {
		str[0] -= 'a' - 'A'
	}
	return string(str)
}
