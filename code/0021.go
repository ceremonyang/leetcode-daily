package code

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := ListNode{Val: -1}
	prev := head
	for {
		if list1.Val < list2.Val {
			prev.Next = list1
			list2 = list2.Next
		} else {
			prev.Next = list2
			list1 = list1.Next
		}
		prev = *prev.Next
		if list1.Next == nil {
			prev.Next = list2
			break
		}
		if list2.Next == nil {
			prev.Next = list1
			break
		}
	}
	return head.Next
}
