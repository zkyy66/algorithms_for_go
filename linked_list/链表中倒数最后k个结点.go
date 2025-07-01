package linked_list

//描述
//输入一个长度为 n 的链表，设链表中的元素的值为 ai ，返回该链表中倒数第k个节点。
//如果该链表长度小于k，请返回一个长度为 0 的链表。
func FindKthToTail(pHead *ListNode, k int) *ListNode {
	// write code here
	if pHead == nil {
		return nil
	}
	if k <= 0 {
		return nil
	}
	prev := &ListNode{}
	prev.Next = pHead
	cur := pHead
	for i := 0; i < k; i++ {
		if cur == nil {
			return nil
		}
		cur = cur.Next
	}
	for cur != nil {
		cur = cur.Next
		prev = prev.Next
	}
	return prev.Next
}
