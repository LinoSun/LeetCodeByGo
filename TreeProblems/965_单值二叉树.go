package TreeProblems

/*如果二叉树每个节点都具有相同的值，那么该二叉树就是单值二叉树。

只有给定的树是单值二叉树时，才返回 true；否则返回 false。

 

示例 1：



输入：[1,1,1,1,1,null,1]
输出：true*/

func isUnivalTree(root *TreeNode) bool {
	queue := []*TreeNode{root}
	num := root.Val

	for len(queue) != 0{
		n := len(queue)
		for i := 0;i<n;i++{
			if queue[i].Val != num{
				return false
			}
			if queue[i].Left != nil{
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil{
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[n:]
	}
	return true
}
