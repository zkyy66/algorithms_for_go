package linked_list

//给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。 例如链表1-》2-3〉4.。。返回4321
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == nil {
		return head
	}
	root := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return root
}
func reverseList204(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == nil {
		return head
	}
	root := reverseList204(head.Next)
	head.Next.Next = head
	head.Next = nil
	return root
}

//
//作者：代码随想录
//链接：https://leetcode.cn/problems/reverse-linked-list/solutions/1718107/by-carlsun-2-uh1f/

//双指针方式
//首先定义一个cur指针，指向头结点，再定义一个pre指针，初始化为null。
//然后就要开始反转了，首先要把 cur->next 节点用tmp指针保存一下，也就是保存一下这个节点。
//为什么要保存一下这个节点呢，因为接下来要改变 cur->next 的指向了，将cur->next 指向pre ，此时已经反转了第一个节点了。
//接下来，就是循环走如下代码逻辑了，继续移动pre和cur指针。
//最后，cur 指针已经指向了null，循环结束，链表也反转完毕了。 此时我们return pre指针就可以了，pre指针就指向了新的头结点。
func reverseListZhizhen(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
