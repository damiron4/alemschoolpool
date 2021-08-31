package piscine

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}
	ans := 0
	for i := 1; start != 1; i++ {
		if start%2 == 0 {
			start /= 2
		} else {
			start = start*3 + 1
		}
		ans = i
	}
	return ans
}
