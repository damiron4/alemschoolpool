package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	temp := l
	if temp == nil {
		return nil
	}
	temp.Next = ListSort(temp.Next)
	// for temp != nil {
	if temp.Next != nil && temp.Data > temp.Next.Data {
		temp = MoveN(temp)
	}
	//temp = temp.Next
	//}
	return temp
}

func MoveN(l *NodeI) *NodeI {
	p := l
	n := l.Next
	s := n
	for n != nil && l.Data > n.Data {
		p = n
		n = n.Next
	}
	p.Next = l
	l.Next = n
	return s
}
