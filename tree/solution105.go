package tree

func buildTree2(preorder []int, inorder []int) *TreeNode {
	index := make(map[int]int)
	for i, val := range inorder {
		index[val] = i
	}
	var traversal func(inorder []int, ii, ij int, preorder []int, pi, pj int) *TreeNode
	traversal = func(inorder []int, ii, ij int, preorder []int, pi, pj int) *TreeNode {
		if pi == pj {
			return nil
		}
		node := &TreeNode{Val: preorder[pi]}
		node.Left = traversal(inorder, ii, index[preorder[pi]], preorder, pi+1, pi+1+index[preorder[pi]]-ii)
		node.Right = traversal(inorder, index[preorder[pi]]+1, ij, preorder, pi+1+index[preorder[pi]]-ii, pj)
		return node
	}
	return traversal(inorder, 0, len(inorder), preorder, 0, len(preorder))
}
