package tree

func pathSum(root *TreeNode, targetSum int) [][]int {
	res := make([][]int, 0)
	list := make([]int, 0)
	if root == nil {
		return res
	}
	var helper func(node *TreeNode, target int)
	helper = func(node *TreeNode, target int) {
		list = append(list, node.Val)

		if node.Left == nil && node.Right == nil {
			if target == node.Val {
				copyList := make([]int, len(list))
				copy(copyList, list)
				res = append(res, copyList)
			}
		}

		if node.Left != nil {
			helper(node.Left, target-node.Val)
		}

		if node.Right != nil {
			helper(node.Right, target-node.Val)
		}

		list = list[0 : len(list)-1]
	}
	helper(root, targetSum)
	return res
}
