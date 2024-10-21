package leetcodehot100

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func reverseListOther(head *ListNode) *ListNode {
	var p *ListNode
	current := head
	p = nil
	for current != nil {
		next := current.Next
		current.Next = p
		p = current
		current = next
	}
	return p
}
