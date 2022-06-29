package list

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	cur := head.Next
	pre := head
	for cur != nil {
		for cur != nil && cur.Val == val {
			cur = cur.Next
		}
		pre.Next = cur
		pre = cur
		if cur != nil {
			cur = cur.Next
		}
	}
	if head.Val == val {
		head = head.Next
	}
	return head
}
