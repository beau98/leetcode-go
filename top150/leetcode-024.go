package top150

func swapPairs(head *ListNode) *ListNode {
	//help node
	sentienl := &ListNode{0, nil}
	sentienl.Next = head
	p := sentienl
	for p.Next != nil && p.Next.Next != nil {
		next1 := p.Next
		next2 := next1.Next
		next1.Next = next2.Next
		next2.Next = next1
		p.Next = next2
		p = p.Next.Next
	}
	return sentienl.Next
}
