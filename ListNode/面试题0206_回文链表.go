package ListNode

func isPalindrome(head *ListNode) bool {
	if head == nil{
		return true
	}
	fast, slow := head, head
	// 找到中间的链表
	stack:=make([]*ListNode,1)
	stack[0] = head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		stack = append(stack, slow)
	}
	if fast.Next == nil {
		stack = stack[:len(stack)-1]
	}
	for slow.Next != nil {
		if slow.Next.Val != stack[len(stack)-1].Val {
			return false
		}
		slow = slow.Next
		stack = stack[:len(stack)-1]
	}
	return len(stack) == 0
}
