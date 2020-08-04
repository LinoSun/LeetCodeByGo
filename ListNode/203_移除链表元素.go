package ListNode


/*删除链表中等于给定值 val 的所有节点。

示例:

输入: 1->2->6->3->4->5->6, val = 6
输出: 1->2->3->4->5
*/

func removeElements(head *ListNode, val int) *ListNode {
    // 最佳实践
    res := &ListNode{0,head}
    pre,cur := res,head

    for cur != nil{
        if cur.Val == val{
            pre.Next = cur.Next
            cur = cur.Next
        } else {
            pre = cur
            cur = cur.Next
        }
    }
    return res.Next
}
