package tree

import "math"

func getMinimumDifference(root *TreeNode) int {
	var pre *TreeNode
	var traversal func(node *TreeNode)
	min := math.MaxInt

	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		if pre != nil && node.Val-pre.Val < min {
			min = node.Val - pre.Val
		}
		pre = node
		traversal(node.Right)
	}

	traversal(root)
	return min
}
