package leetcode

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	tempnode := head
	for tempnode.Next != nil {
		if tempnode.Val != tempnode.Next.Val {
			tempnode = tempnode.Next
			continue
		}
		tempnode.Next = tempnode.Next.Next
	}
	return head
}
