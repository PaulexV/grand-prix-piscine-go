package piscine

func IsAlpha(str string) bool {

	for _, letter := range str {
		if !((letter >= 'A' && letter <= 'Z') || (letter >= 'a' && letter <= 'z') || (letter >= '0' && letter <= '9')) {
			return false
		}
	}
	return true
}
