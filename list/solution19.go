package list

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	pre := dummy
	slow, fast := head, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for slow != nil && fast != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next
	}
	pre.Next = slow.Next
	return dummy.Next
}
