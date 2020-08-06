package ListNode

/*反转一个单链表。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
进阶:
你可以迭代或递归地反转链表。你能否用两种方法解决这道题？*/

func reverseList(head *ListNode) *ListNode {
	// 设置哨兵链表
	var sentinel *ListNode = nil
	cur := head
	// 每一次的归位方向很重要
	for cur != nil{
		tmp := cur.Next
		cur.Next = sentinel
		sentinel = cur
		cur = tmp
	}
	return sentinel
}
