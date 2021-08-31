package piscine

func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	for root != nil {
		root = root.Left
		if root.Left == nil {
			return root
		}
	}
	return nil
}
