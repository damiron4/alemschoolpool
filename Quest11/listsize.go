package piscine

func ListSize(l *List) int {
	var count int = 0
	for l.Head != nil {
		count++
		l.Head = l.Head.Next
	}
	return count
}
