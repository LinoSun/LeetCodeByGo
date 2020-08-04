package ListNode

/*将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 

 

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4*/

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 哨兵站点设立很关键
	newL:= &ListNode{}
	res := newL
	for l1 != nil && l2 != nil{
		if l1.Val < l2.Val{
			newL.Next = l1
			l1 = l1.Next
		} else {
			newL.Next = l2
			l2 = l2.Next
		}
		newL = newL.Next
	}
	if l1 == nil{
		newL.Next = l2
	}
	if l2 == nil{
		newL.Next = l1
	}
	return res.Next
}
