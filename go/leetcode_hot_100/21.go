package leetcodehot100

import ()

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy := &ListNode{}
    current := dummy
    for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		}else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}
	if list1 != nil {
		current.Next = list1
	}else {
		current.Next = list2
	}

	return dummy.Next
}

func mergeTwoListsOther(list1 *ListNode,list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoListsOther(list1.Next,list2)
		return list1
	}
	list2.Next = mergeTwoListsOther(list1,list2.Next)
	return list2
}