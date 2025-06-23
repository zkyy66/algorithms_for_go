package linked_list

func removeElement(head *ListNode, val int) *ListNode {
	points := &ListNode{}
	points.Next = head
	cur := points
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return points.Next
}
