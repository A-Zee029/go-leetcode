package others

import "container/list"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		l := queue.Len()
		var pre *Node
		for i := 0; i < l; i++ {
			node := queue.Remove(queue.Front()).(*Node)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			if pre != nil {
				pre.Next = node
			}
			pre = node
		}
	}
	return root
}
