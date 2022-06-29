package tree

import (
	"container/list"
	"math"
)

func largestValues(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		max := math.MinInt
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			if max < node.Val {
				max = node.Val
			}
		}
		res = append(res, max)
	}
	return res
}
