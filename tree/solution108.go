package tree

func sortedArrayToBST(nums []int) *TreeNode {
	var helper func(nums []int, l, r int) *TreeNode
	helper = func(nums []int, l, r int) *TreeNode {
		if l >= r {
			return nil
		}
		mid := (r-l-1)/2 + l
		node := &TreeNode{Val: nums[mid]}
		node.Left = helper(nums, l, mid)
		node.Right = helper(nums, mid+1, r)
		return node
	}
	return helper(nums, 0, len(nums))
}
