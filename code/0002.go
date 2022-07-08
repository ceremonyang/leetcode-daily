package code

/*
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func EndInsert(val int, head *ListNode) *ListNode {
	node := &ListNode{
		Val:  val,
		Next: nil,
	}
	if head.Next == nil {
		head.Next = node
		return head
	}
	EndInsert(val, head.Next)
	return head

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l := &ListNode{}
	carry := 0
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + carry
		carry = sum / 10
		l = EndInsert(sum%10, l)
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		sum := l1.Val + carry
		carry = sum / 10
		l = EndInsert(sum%10, l)
		l1 = l1.Next
	}
	for l2 != nil {
		sum := l2.Val + carry
		carry = sum / 10
		l = EndInsert(sum%10, l)
		l2 = l2.Next
	}
	if carry != 0 {
		l = EndInsert(carry, l)
	}

	return l.Next
}
