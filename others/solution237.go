package others

func deleteNode(node *ListNode) {
	next := node.Next
	node.Val = next.Val
	node.Next = next.Next
}
