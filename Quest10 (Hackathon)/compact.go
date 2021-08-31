package piscine

const N = 6

func Compact(ptr *[]string) int {
	ans := len(*ptr)
	for _, word := range *ptr {
		if word == "" {
			ans--
		}
	}
	index := 0
	arr := make([]string, ans)
	for _, word := range *ptr {
		if word != "" {
			arr[index] = word
			index++
		}
	}
	for i := 0; i < len(arr); i++ {
		(*ptr)[i] = arr[i]
	}
	*ptr = (*ptr)[0:ans]
	return ans
}
