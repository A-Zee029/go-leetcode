package tree

func kthSmallest(root *TreeNode, k int) int {
	cnt := 0
	res := 0
	inOrder(root, k, &cnt, &res)
	return res
}

func inOrder(node *TreeNode, k int, cnt *int, res *int) {
	if node == nil {
		return
	}
	inOrder(node.Left, k, cnt, res)
	*cnt++
	if *cnt == k {
		*res = node.Val
		return
	}
	inOrder(node.Right, k, cnt, res)
}
