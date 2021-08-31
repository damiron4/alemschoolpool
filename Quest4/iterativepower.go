package piscine

func IterativePower(nb int, power int) int {
	if power == 0 {
		return 1
	} else if nb == 0 || power < 0 {
		return 0
	}
	temp := nb
	for i := 1; i < power; i++ {
		nb *= temp
	}
	return nb
}
