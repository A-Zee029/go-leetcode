package tree

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor2(root.Left, p, q)

	} else if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor2(root.Right, p, q)
	} else {
		return root
	}
}
