package tree

func minCameraCover(root *TreeNode) int {
	result := 0

	var traversal func(node *TreeNode) int
	// 0-未覆盖 1-有摄像头 2-有覆盖
	traversal = func(node *TreeNode) int {
		if node == nil {
			return 2
		}
		left := traversal(node.Left)
		right := traversal(node.Right)
		if left == 2 && right == 2 {
			return 0
		} else if left == 0 || right == 0 {
			result++
			return 1
		} else {
			return 2
		}
	}
	if traversal(root) == 0 {
		result++
	}
	return result
}
