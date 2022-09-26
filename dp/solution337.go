package dp

import (
	"leetcode/tree"
)

func rob(root *tree.TreeNode) int {
	ints := robTree(root)
	if ints[0] > ints[1] {
		return ints[0]
	} else {
		return ints[1]
	}
}

func robTree(root *tree.TreeNode) []int {
	res := make([]int, 2)
	if root == nil {
		res[0], res[1] = 0, 0
		return res
	}
	left := robTree(root.Left)
	right := robTree(root.Right)
	res[1] = root.Val + left[0] + right[0]
	leftMax := left[0]
	if left[1] > leftMax {
		leftMax = left[1]
	}
	rightMax := right[0]
	if right[1] > rightMax {
		rightMax = right[1]
	}
	res[0] = leftMax + rightMax
	return res
}
