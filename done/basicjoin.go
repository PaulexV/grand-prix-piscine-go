package piscine

func BasicJoin(strs []string) string {

	var res string

	for index := range strs {
		res += strs[index]
	}
	return res
}
