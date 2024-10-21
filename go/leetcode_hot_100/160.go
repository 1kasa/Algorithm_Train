package leetcodehot100

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	mp := map[*ListNode]bool{}
	for index := headA; index != nil; index = index.Next {
		mp[index] = true
	}

	for index := headB; index != nil; index = index.Next {
		if mp[index] {
			return index
		}
	}

	return nil
}

// 双指针
func getIntersectionNodeOther(headA, headB *ListNode) *ListNode {
	if headA == nil && headB == nil {
		return nil
	}

	pointA, pointB := headA, headB
	for pointA != pointB {
		if pointA == nil {
			pointA = pointB
		} else {
			pointA = pointA.Next
		}
		if pointB == nil {
			pointB = pointA
		} else {
			pointB = pointB.Next
		}
	}
	return pointA
}
