package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummyHead := &ListNode{0, nil}
	dummyHead.Next = head
	cur := dummyHead
	for cur.Next != nil && cur.Next.Next != nil {
		node1 := cur.Next
		node2 := cur.Next.Next.Next

		cur.Next = cur.Next.Next
		cur.Next.Next = node1
		cur.Next.Next.Next = node2
		cur = cur.Next.Next
	}
	return dummyHead.Next
}

func swapPairsOther(head *ListNode) *ListNode {
	if head == nil || head.Next ==nil {
       return head
	}
	next := head.Next
	head.Next = swapPairsOther(next.Next)
	next.Next = head
	return next
}