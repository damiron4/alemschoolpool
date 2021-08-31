package piscine

func DeleteNode(root *TreeNode, key string) *TreeNode {
	if root == nil {
		return root
	}
	if key < root.Data {
		root.Left = DeleteNode(root.Left, key)
	} else if key > root.Data {
		root.Right = DeleteNode(root.Right, key)
	} else {
		if root.Left == nil {
			temp := root.Right
			return temp
		} else if root.Right == nil {
			temp := root.Left
			return temp
		}
		temp := BTreeMin(root.Right)
		root.Data = temp.Data
		root.Right = DeleteNode(root.Right, key)
	}
	return root
}

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	root = DeleteNode(root, node.Data)
	root = BTreeInsertData(root, rplc.Data)
	return root
}
