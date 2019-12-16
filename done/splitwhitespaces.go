package piscine

func SplitWhiteSpaces(str string) []string {

	size := 1
	i_res := 0

	for i, char := range str {
		if (char == ' ' && str[i+1] != ' ') || char == '\n' || (char == '\t' && str[i+1] != '\t') {
			size++
		}
	}
	res := make([]string, size)
	word := ""
	for i, char := range str {
		if (char == ' ' && str[i+1] != ' ') || char == '\n' || (char == '\t' && str[i+1] != '\t') {
			res[i_res] = word
			i_res++
			word = ""
		} else if char != ' ' || char != '\n' || char != '\t' {
			if !(str[i] == ' ') {
				word += string(char)
			}
		}
	}
	res[i_res] = word
	return res
}
