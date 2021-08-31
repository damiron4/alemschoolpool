package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	count := 0
	for l != nil {
		if pos == count {
			return l
		}
		l = l.Next
		count++
	}
	return nil
}
