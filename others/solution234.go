package others

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	arr := make([]int, 0)
	size := 0
	for ; head != nil; head = head.Next {
		arr = append(arr, head.Val)
		size++
	}
	fmt.Println(arr)
	fmt.Println(size)
	for i := 0; i < size/2; i++ {
		if arr[i] != arr[size-i-1] {
			return false
		}
	}
	return true
}
