package piscine

func f(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	} else {
		return -1
	}
}

func IsSorted(f func(a, b int) int, a []int) bool {
	indp := true
	indn := true
	for i := 1; i < len(a); i++ {
		arg := f(a[i-1], a[i])
		if arg < 0 {
			indp = false
		}
	}
	for i := 1; i < len(a); i++ {
		arg := f(a[i-1], a[i])
		if arg > 0 {
			indn = false
		}
	}
	return indp || indn
}
