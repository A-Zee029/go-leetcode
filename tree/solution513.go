package tree

import "container/list"

func findBottomLeftValue(root *TreeNode) int {
	queue := list.New()
	queue.PushBack(root)
	res := 0
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			if i == 0 {
				res = node.Val
			}
		}
	}
	return res
}
