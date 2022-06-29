package tree

func findMode(root *TreeNode) []int {
	var pre *TreeNode
	var traversal func(node *TreeNode)
	cnt := 0
	max := 0
	res := make([]int, 0)

	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}

		traversal(node.Left)
		if pre == nil || pre.Val != node.Val {
			cnt = 1
		} else {
			cnt++
		}
		pre = node

		if cnt == max {
			res = append(res, node.Val)
		}
		if cnt > max {
			max = cnt
			res = res[len(res):]
			res = append(res, node.Val)
		}
		traversal(node.Right)
	}

	traversal(root)
	return res

}
