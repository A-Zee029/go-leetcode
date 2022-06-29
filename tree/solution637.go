package tree

import "container/list"

func averageOfLevels(root *TreeNode) []float64 {
	res := make([]float64, 0)
	if root == nil {
		return res
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		sum := 0.0
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			sum += float64(node.Val)
		}
		sum /= float64(length)
		res = append(res, sum)
	}
	return res
}
