package piscine

func IsUpper(str string) bool {

	for _, letter := range str {
		if !(letter >= 'A' && letter <= 'Z') {
			return false
		}
	}
	return true
}
