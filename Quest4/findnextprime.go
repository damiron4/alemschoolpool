package piscine

func FindNextPrime(nb int) int {
	if nb >= 2147483647 {
		return 0
	} else if nb <= 2 {
		return 2
	}
	count := 0
	for i := nb; true; i++ {
		count = 0
		for j := 2; j <= i/2; j++ {
			if i%j == 0 {
				j = i/2 + 1
				continue
			}
			count++
		}
		if count == (i-2)/2 {
			return i
		}
	}
	return 0
}
