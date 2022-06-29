package list

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{}
	dummy.Next = head
	pre := dummy
	for head != nil && head.Next != nil {
		next1 := head.Next
		next2 := next1.Next
		head.Next = next2
		next1.Next = head
		pre.Next = next1
		pre = head
		head = next2
	}
	return dummy.Next
}
