package ListNode

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
//	var newL =  *ListNode{0}
//	return newL
//}

func min(x,y int)int{
	if x < y {
		return x
	}
	return y
}
