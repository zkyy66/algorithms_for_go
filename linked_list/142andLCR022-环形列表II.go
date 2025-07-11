package linked_list

/**
给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。

不允许修改 链表
*/
//借助于141题判断是否有环的思路，通过mapHash方式
func detectCycleForMap(head *ListNode) *ListNode {
	isCycle := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := isCycle[head]; ok {
			return head
		}
		isCycle[head] = struct{}{}
		head = head.Next
	}
	return nil
}

func detectCycleForZhiZhen(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil || fast.Next != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if fast == slow {
			p := head
			for slow != p {
				slow = slow.Next
				p = p.Next
			}
			return p //return slow 结果一样？
		}
	}
	return nil
}

func EntryNodeOfLoop(pHead *ListNode) *ListNode {
	if pHead == nil || pHead.Next == nil {
		return nil
	}
	// 快慢指针
	fast := pHead
	last := pHead
	
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		last = last.Next
		if fast == last {
			break
		}
	}
	// 第一次相遇
	if last == nil || fast == nil {
		// 无环
		return nil
	}
	last = pHead
	for last != fast {
		last = last.Next
		fast = fast.Next
	}
	return last
}
