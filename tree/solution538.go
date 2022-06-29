package tree

func convertBST(root *TreeNode) *TreeNode {

	sum := 0
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Right)
		node.Val += sum
		sum = node.Val
		traversal(node.Left)
	}

	traversal(root)

	return root
}
