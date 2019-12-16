package piscine

func IterativePower(nb int, power int) int {

	res := 1

	if power < 0 {
		return 0
	}
	if power == 1 {
		return nb
	}
	for i := 0; i < power; i++ {
		res = res * nb
	}
	return res
}
