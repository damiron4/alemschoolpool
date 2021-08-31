package piscine

func ListReverse(l *List) {
	var prev, next *NodeL
	prev = nil
	next = nil
	current := l.Head
	l.Tail = current
	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	l.Head = prev
}
