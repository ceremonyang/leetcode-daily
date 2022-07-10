package code

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head = &ListNode{Val: -1}
	last := head
	for list1.Next == nil && list2.Next != nil {
		if list1.Val > list2.Val {
			last.Next = list2
			list2 = list2.Next
		} else {
			last.Next = list1
			list1 = list1.Next
		}
		last = last.Next
	}
	return head.Next
}
