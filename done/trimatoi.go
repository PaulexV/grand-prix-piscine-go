package piscine

func TrimAtoi(s string) int {
	str := []byte(s)
	var neg bool = false
	var result string
	for _, i := range str {
		if rune(i) >= 48 && rune(i) <= 57 {
			result += string(i)
		}
		if rune(i) == 45 && result == "" {
			neg = true
		}
	}
	if neg == false {
		return basicAtoi(result)
	} else {
		return basicAtoi(result) * -1
	}
}

func basicAtoi(s string) int {
	rep := 0
	for _, index := range s {
		nmb := 0
		for i := '1'; i <= index; i++ {
			nmb++
		}
		rep = (rep * 10) + nmb
	}
	return rep
}
