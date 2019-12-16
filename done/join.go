package piscine

func Join(strs []string, sep string) string {

	lenS := 0
	str := ""

	for range strs {
		lenS++
	}
	for i, value := range strs {
		if i != lenS-1 {
			str += value + sep
		} else {
			str += value
		}
	}
	return str
}
