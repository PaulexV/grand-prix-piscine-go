package piscine

func UltimateDivMod(a, b *int) {

	c := *a
	*a = *a / *b
	*b = c % *b
}
