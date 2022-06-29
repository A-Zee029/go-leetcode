package tree

func isSymmetric(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return true
	}
	return helper(root.Left, root.Right)
}

func helper(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil || left.Val != right.Val {
		return false
	}
	return helper(left.Left, right.Right) && helper(right.Left, left.Right)
}
