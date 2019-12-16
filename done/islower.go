package piscine

func IsLower(str string) bool {

	for _, letter := range str {
		if !(letter >= 'a' && letter <= 'z') {
			return false
		}
	}
	return true
}
