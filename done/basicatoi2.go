package piscine

func BasicAtoi2(s string) int {
	res := 0
	for _, val := range s {
		a := 0
		if val >= '0' && val <= '9' {

			for i := '1'; i <= val; i++ {

				a++
			}
			res = res*10 + a
		} else {
			return 0
		}
	}
	return res
}
