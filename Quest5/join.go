package piscine

func Join(strs []string, sep string) string {
	var str string
	for i := 0; i < len(strs); i++ {
		str += strs[i]
		if i != len(strs)-1 {
			str = str + sep
		}
	}
	return str
}
