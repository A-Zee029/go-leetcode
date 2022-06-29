package tree

func buildTree(inorder []int, postorder []int) *TreeNode {
	index := make(map[int]int)
	for i, val := range inorder {
		index[val] = i
	}

	var traversal func(inorder []int, ii, ij int, postorder []int, pi, pj int) *TreeNode
	traversal = func(inorder []int, ii, ij int, postorder []int, pi, pj int) *TreeNode {
		if pi == pj {
			return nil
		}
		node := &TreeNode{Val: postorder[pj-1]}
		node.Left = traversal(inorder, ii, index[postorder[pj-1]], postorder, pi, pi+index[postorder[pj-1]]-ii)
		node.Right = traversal(inorder, index[postorder[pj-1]]+1, ij, postorder, pi+index[postorder[pj-1]]-ii, pj-1)
		return node
	}
	return traversal(inorder, 0, len(inorder), postorder, 0, len(postorder))
}
