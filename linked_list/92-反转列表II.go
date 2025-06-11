package linked_list

//给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表
//输入：head = [1,2,3,4,5], left = 2, right = 4
//输出：[1,4,3,2,5]

//使用三个指针变量 pre、curr、next 来记录反转的过程中需要的变量，它们的意义如下：
//
//curr：指向待反转区域的第一个节点 left；
//next：永远指向 curr 的下一个节点，循环过程中，curr 变化以后 next 会变化；
//pre：永远指向待反转区域的第一个节点 left 的前一个节点，在循环过程中不变。
//
//作者：力扣官方题解
//链接：https://leetcode.cn/problems/reverse-linked-list-ii/solutions/634701/fan-zhuan-lian-biao-ii-by-leetcode-solut-teyq/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	tempNode := &ListNode{Val: -1}
	tempNode.Next = head
	pre := tempNode
	for i := 0; i <= left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return tempNode.Next
}
