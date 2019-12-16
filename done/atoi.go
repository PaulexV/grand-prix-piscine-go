package piscine

func Atoi(s string) int {

	res := 0
	isNeg := 0

	for _, val := range s {
		a := 0
		if val == '-' {

			isNeg = 1
		}
		if (val >= '0' && val <= '9') || val == '-' || val == '+' {

			for i := '1'; i <= val; i++ {

				a++
			}
			res = res*10 + a
		} else if isNeg == 1 {

			return -res
		} else {

			return 0
		}
	}
	return res
}
