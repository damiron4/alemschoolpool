package piscine

func Max(a []int) int {
	max := -2147483648
	if len(a) == 0 {
		return 0
	}
	for _, val := range a {
		if max < val {
			max = val
		}
	}
	return max
}
