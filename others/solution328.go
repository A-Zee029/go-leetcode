package others

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}
	var second *ListNode
	odd := head
	for odd != nil && odd.Next != nil {
		even := odd.Next
		if second == nil {
			second = even
		}
		odd.Next = even.Next
		if odd.Next == nil {
			break
		}
		odd = odd.Next
		even.Next = odd.Next
	}
	odd.Next = second
	return head
}
