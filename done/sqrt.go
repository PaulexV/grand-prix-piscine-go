package piscine

func Sqrt(nb int) int {

	if nb == 1 {
		return 1
	}
	for racine := 1; racine <= nb/2; {
		sqrt := racine * racine
		if sqrt == nb {
			return racine
		}
		racine += 1
	}
	return 0
}
