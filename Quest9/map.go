package piscine

func Map(f func(int) bool, a []int) []bool {
	arr := make([]bool, len(a))
	for index, ans := range a {
		arr[index] = f(ans)
	}
	return arr
}
