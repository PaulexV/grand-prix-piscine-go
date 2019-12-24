package piscine

func IsSorted(f func(a, b int) int, tab []int) bool {

	pas := true
	logique := true

	for i := range tab {
		if i > 0 {
			if f(tab[i-1], tab[i]) < 0 {
				pas = false
			}
			if f(tab[i-1], tab[i]) > 0 {
				logique = false
			}
		}
	}
	if pas || logique {
		return true
	}
	return false
}
