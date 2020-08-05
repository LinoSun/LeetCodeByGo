package ListNode

/*给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。

进阶：

你能尝试使用一趟扫描实现吗？*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 快慢指针 最佳实践
	sentinel := new(ListNode)
	sentinel.Next = head

	preN := sentinel
	i:=1
	for head != nil{
		if i >= n{
			// 慢指针
			preN = preN.Next
		}
		// 快指针
		head = head.Next
		i++
	}
	preN.Next = preN.Next.Next
	return sentinel.Next
}
