package tree

import "strconv"

func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)
	getPath(root, "", &res)
	return res
}

func getPath(root *TreeNode, s string, res *[]string) {
	if root == nil {
		return
	}
	if s != "" {
		s += "->"
	}
	s += strconv.Itoa(root.Val)

	if root.Left == nil && root.Right == nil {
		*res = append(*res, s)
	}
	getPath(root.Left, s, res)
	getPath(root.Right, s, res)
}
