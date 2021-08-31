package piscine

func ActiveBits(n int) int {
	ans := 1
	for i := 0; n != 1; i++ {
		if n%2 == 1 {
			ans++
		}
		n /= 2
	}
	return ans
}
