package tree

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftVal, rightVal := sumOfLeftLeaves(root.Left), sumOfLeftLeaves(root.Right)
	midVal := 0
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		midVal = root.Left.Val
	}
	return leftVal + rightVal + midVal
}
