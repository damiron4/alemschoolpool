package piscine

func Index(s string, toFind string) int {
	for i := 0; i < len(s); i++ {
		if s[i] == toFind[0] {
			count := 1
			for j := 1; j < len(toFind) && i+j < len(s); j++ {
				if s[i+j] != toFind[j] {
					break
				} else {
					count++
				}
			}
			if count == len(toFind) {
				return i
			}
		}
	}
	return -1
}
