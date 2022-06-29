package tree

func constructMaximumBinaryTree(nums []int) *TreeNode {
	var traversal func(nums []int, i, j int) *TreeNode
	traversal = func(nums []int, i, j int) *TreeNode {
		if i == j {
			return nil
		}
		max := i
		for idx := i; idx < j; idx++ {
			if nums[idx] > nums[max] {
				max = idx
			}
		}
		node := &TreeNode{Val: nums[max]}
		node.Left = traversal(nums, i, max)
		node.Right = traversal(nums, max+1, j)
		return node
	}
	return traversal(nums, 0, len(nums))
}
