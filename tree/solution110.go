package tree

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if isBalanced(root.Left) && isBalanced(root.Right) {
		dif := getHeight(root.Left) - getHeight(root.Right)
		if dif < 0 {
			dif = -dif
		}
		return dif <= 1
	}
	return false
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lHeight, rHeight := getHeight(root.Left), getHeight(root.Right)
	if lHeight < rHeight {
		return 1 + rHeight
	} else {
		return 1 + lHeight
	}
}
