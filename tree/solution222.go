package tree

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lHeight, rHeight := 0, 0
	left, right := root.Left, root.Right
	for left != nil {
		left = left.Left
		lHeight++
	}
	for right != nil {
		right = right.Right
		rHeight++
	}
	if lHeight == rHeight {
		return 2<<lHeight - 1
	} else {
		return 1 + countNodes(root.Left) + countNodes(root.Right)
	}
}
