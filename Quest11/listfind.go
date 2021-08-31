package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	smth := l.Head
	if smth == nil {
		return nil
	}
	for smth != nil {
		if comp(smth.Data, ref) {
			break
		}
		smth = smth.Next
	}
	return &smth.Data
}
