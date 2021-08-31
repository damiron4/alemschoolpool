package piscine

func RecursivePower(nb int, power int) int {
	if power == 0 {
		return 1
	} else if nb == 0 || power < 0 {
		return 0
	}
	return RecursivePower(nb, power-1) * nb
}
