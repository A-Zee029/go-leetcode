package tree

func deleteNode2(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == key {
		newRoot := root.Right
		if newRoot != nil {
			node := newRoot
			for node.Left != nil {
				node = node.Left
			}
			node.Left = root.Left
		}
		if newRoot == nil {
			newRoot = root.Left
		}
		return newRoot
	}
	if root.Val < key {
		root.Right = deleteNode2(root.Right, key)
	}
	if root.Val > key {
		root.Left = deleteNode2(root.Left, key)
	}
	return root
}
