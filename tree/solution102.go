package tree

import "container/list"

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil { //防止为空
		return res
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		tmp := make([]int, length)
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			tmp[i] = node.Val
		}
		res = append(res, tmp)
	}
	return res
}
