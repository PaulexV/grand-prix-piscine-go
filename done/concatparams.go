package piscine

func ConcatParams(args []string) string {

	var str string
	lenS := 0

	for range args {
		lenS++
	}
	for i := range args {
		if i != lenS-1 {
			str += args[i] + "\n"
		} else {
			str += args[i]
		}
	}
	return str
}
