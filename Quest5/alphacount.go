package piscine

func AlphaCount(s string) int {
	count := 0
	slice := []byte(s)
	for i := 0; i < len(s); i++ {
		if (slice[i] >= 65 && slice[i] <= 90) || (slice[i] >= 97 && slice[i] <= 122) {
			count++
		}
	}
	return count
}
