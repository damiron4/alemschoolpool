package piscine

func Swap(a *int, b *int) {
	var c int = *a
	*a = *b
	*b = c
}
