package piscine

func ConcatParams(args []string) string {
	var str string
	spaces := make([]string, len(args)-1)
	for i := 0; i < len(args)-1; i++ {
		spaces[i] += string("\n")
	}
	for i := 0; i < len(args); i++ {
		str += args[i]
		if i != len(args)-1 {
			str += spaces[i]
		}
	}
	return str
}
