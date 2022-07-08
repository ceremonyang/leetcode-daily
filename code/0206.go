package code

func reverseList(head *ListNode) *ListNode {
	node := ListNode{}
	if head == nil {
		node.Next = head
	}
	reverseList(head.Next)
	return node.Next
}
