package piscine

func ListMerge(l1 *List, l2 *List) {
	if l1.Head == nil {
		l1.Head, l1.Tail = l2.Head, l2.Tail
		return
	}
	var temp, temp2 *NodeL
	temp = l1.Head
	temp2 = temp
	for temp2.Next != nil {
		temp2 = temp2.Next
	}
	temp2.Next = l2.Head
}
