package tree

import "container/list"

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := list.New()
	queue.PushBack(root)
	res := 0
	for queue.Len() > 0 {
		l := queue.Len()
		for ; l > 0; l-- {
			front := queue.Remove(queue.Front()).(*TreeNode)
			if front.Left != nil {
				queue.PushBack(front.Left)
			}
			if front.Right != nil {
				queue.PushBack(front.Right)
			}
		}
		res++
	}
	return res
}
