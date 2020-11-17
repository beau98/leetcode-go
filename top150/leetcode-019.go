package top150

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

/**
Given the head of a linked list, remove the nth node from the end of the list and return its head.

Follow up: Could you do this in one pass?
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	newHead := ListNode{0, head}
	p1 := &newHead
	p2 := &newHead

	for i := 0; i < n+1; i++ {
		p1 = p1.Next
	}
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	p2.Next = p2.Next.Next
	return newHead.Next
}
