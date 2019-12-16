package piscine

func is_prime(nb int) bool {

	if nb < 2 {
		return false
	}
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(nb int) int {

	for is_prime(nb) == false {
		nb++
	}
	return nb
}
