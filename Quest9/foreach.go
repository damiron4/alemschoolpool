package piscine

func ForEach(f func(int), a []int) {
	for _, ans := range a {
		f(ans)
	}
}
