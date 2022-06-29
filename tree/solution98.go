package tree

func isValidBST(root *TreeNode) bool {

	var traversal func(node *TreeNode) (bool, int, int)

	traversal = func(node *TreeNode) (bool, int, int) {
		min, max := node.Val, node.Val
		if node.Left != nil {
			left, lmin, lmax := traversal(node.Left)
			if !left || lmax >= node.Val {
				return false, lmin, lmax
			}
			min = lmin
		}

		if node.Right != nil {
			right, rmin, rmax := traversal(node.Right)
			if !right || rmin <= node.Val {
				return false, rmin, rmax
			}
			max = rmax
		}
		return true, min, max
	}

	res, _, _ := traversal(root)
	return res
}
