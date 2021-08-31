package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	if n1 == nil {
		return nil
	}
	var temp, temp2 *NodeI
	temp = n1
	temp2 = temp
	for temp2.Next != nil {
		temp2 = temp2.Next
	}
	temp2.Next = n2
	return ListSort(n1)
}
