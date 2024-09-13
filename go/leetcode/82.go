package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicatesV2(head *ListNode) *ListNode {
	cur, prev := head, head
	for cur != nil && cur.Next != nil {
		if cur.Val != cur.Next.Val {
			prev = cur
			cur = cur.Next
			continue
		}

		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
		}

		if prev.Val == cur.Val {
			prev = cur.Next
			head = prev
			cur = cur.Next
			continue
		}
		prev.Next = cur.Next
		cur = cur.Next 
	}
	return head
}

func deleteDuplicatesV3(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

	if head.Val != head.Next.Val {
		head.Next = deleteDuplicatesV3(head.Next)
		return head
	}

	move := head.Next
	for move != nil && move.Val == head.Val {
		move = move.Next
	}

	return deleteDuplicatesV3(move)
}