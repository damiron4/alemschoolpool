package piscine

func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	for root != nil {
		root = root.Right
		if root.Right == nil {
			return root
		}
	}
	return nil
}
