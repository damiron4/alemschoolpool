package piscine

func Sqrt(nb int) int {
	if nb == 1 {
		return nb
	}
	for i := 0; i < nb; i++ {
		if i*i == nb {
			return i
		}
	}
	return 0
}
