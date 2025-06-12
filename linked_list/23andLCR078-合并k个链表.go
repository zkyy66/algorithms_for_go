package linked_list

//递归做法
/**
暴力做法是，按照 21. 合并两个有序链表 的 题解思路，先合并前两个链表，再把得到的新链表和第三个链表合并，再和第四个链表合并，依此类推。

但是这种做法，平均每个节点会参与到 O(m) 次合并中（用 (1+2+⋯+m)/m 粗略估计），所以总的时间复杂度为 O(L⋅m)。

一个巧妙的思路是，把 lists 一分为二（尽量均分），先合并前一半的链表，再合并后一半的链表，然后把这两个链表合并成最终的链表。如何合并前一半的链表呢？我们可以继续一分为二。如此分下去直到只有一个链表，此时无需合并。

我们可以写一个递归来完成上述逻辑，如果你对递归头晕，请看【基础算法精讲 09】。

按照一分为二再合并的逻辑，递归像是在后序遍历一棵平衡二叉树。由于平衡树的高度是 O(logm)，所以每个链表节点只会出现在 O(logm) 次合并中！这样就做到了更快的 O(Llogm) 时间。
*/
func mergeKLists(lists []*ListNode) *ListNode {
	l := len(lists)
	if l == 0 {
		return nil
	}
	if l == 1 {
		return lists[0]
	}
	left := mergeKLists(lists[:l/2])
	right := mergeKLists(lists[l/2:])
	return mergeTwoLists2(left, right)
}
func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
	if list1 != nil {
		cur.Next = list1
	} else if list2 != nil {
		cur.Next = list2
	}
	return dummy.Next
}

//迭代做法
/***
直接自底向上合并链表：

两两合并：把 lists[0] 和 lists[1] 合并，合并后的链表保存在 lists[0] 中；把 lists[2] 和 lists[3] 合并，合并后的链表保存在 lists[2] 中；依此类推。
四四合并：把 lists[0] 和 lists[2] 合并（相当于合并前四条链表），合并后的链表保存在 lists[0] 中；把 lists[4] 和 lists[6] 合并，合并后的链表保存在 lists[4] 中；依此类推。
八八合并：把 lists[0] 和 lists[4] 合并（相当于合并前八条链表），合并后的链表保存在 lists[0] 中；把 lists[8] 和 lists[12] 合并，合并后的链表保存在 lists[8] 中；依此类推。
依此类推，直到所有链表都合并到 lists[0] 中。最后返回 lists[0]。
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func mergeKLists2(lists []*ListNode) *ListNode {
	m := len(lists)
	if m == 0 {
		return nil
	}
	for step := 1; step < m; step *= 2 {
		for i := 0; i < m-step; i += step * 2 {
			lists[i] = mergeTwoLists2(lists[i], lists[i+step])
		}
	}
	return lists[0]
}
