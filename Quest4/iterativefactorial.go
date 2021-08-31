package piscine

func IterativeFactorial(nb int) int {
	ans := 0
	if nb >= 21 || nb < 0 {
		return ans
	} else if nb == 0 {
		ans = 1
		return ans
	} else {
		ans = 1
	}
	for i := 1; i < nb+1; i++ {
		ans *= i
		if ans < 0 {
			ans = 0
			return ans
		}
	}
	return ans
}
