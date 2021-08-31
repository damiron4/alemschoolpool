package piscine

func UltimateDivMod(a *int, b *int) {
	var c int = *a
	*a = *a / *b
	*b = c % *b
}
