package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	prev := l.Head
	smth := l.Head
	for smth != nil && smth.Data == data_ref {
		l.Head = smth.Next
		smth = l.Head
	}
	for smth != nil {
		if smth.Data != data_ref {
			prev = smth
		}
		prev.Next = smth.Next
		smth = prev.Next
	}
}
