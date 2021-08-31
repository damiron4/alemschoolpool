package piscine

func MakeRange(min, max int) []int {
	a := max - min
	var arr2 []int
	if a <= 0 {
		return arr2
	}
	arr := make([]int, a)
	for i, j := min, 0; i < max; i, j = i+1, j+1 {
		arr[j] += i
	}
	return arr
}
