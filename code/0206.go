package code

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func iReverseList(head *ListNode) *ListNode {
	var node *ListNode
	for head != nil {
		next := head.Next
		head.Next = node
		node = head
		head = next
	}
	return node
}
