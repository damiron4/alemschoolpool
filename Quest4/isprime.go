package piscine

func IsPrime(nb int) bool {
	if nb <= 1 || nb >= 2147483647 {
		return false
	}
	for i := 2; i < nb; i++ {
		if nb%i == 0 && i != nb {
			return false
		}
	}
	return true
}
