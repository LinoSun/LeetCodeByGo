package ListNode

//编写一个程序，找到两个单链表相交的起始节点。

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA,curB := headA,headB
	for curA != curB{
		if curA == nil{
			curA = headB
		} else {
			curA = curA.Next
		}
		if curB == nil{
			curB = headA
		} else {
			curB = curB.Next
		}
	}
	return curA
}
