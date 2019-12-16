package piscine

func IsNumeric(str string) bool {

	for _, letter := range str {
		if !(letter >= '0' && letter <= '9') {
			return false
		}
	}
	return true
}
