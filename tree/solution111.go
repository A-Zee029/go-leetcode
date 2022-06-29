package tree

import "container/list"

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := list.New()
	queue.PushBack(root)
	res := 0
	for queue.Len() > 0 {
		l := queue.Len()
		res++
		for ; l > 0; l-- {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left == nil && node.Right == nil {
				break
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		if l > 0 {
			break
		}
	}
	return res
}
