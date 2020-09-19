package TreeProblems
/*计算给定二叉树的所有左叶子之和。

示例：

3
/ \
9  20
/  \
15   7

在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24
*/

func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	if root == nil{
		return sum
	}
	stack := []*TreeNode{root}
	for len(stack) != 0{
		n := len(stack)
		for i:=0;i<n;i++{
			if stack[i].Left != nil{
				if isLeafNode(stack[i].Left){
					sum += stack[i].Left.Val
				} else {
					stack = append(stack, stack[i].Left)
				}
			}
			if stack[i].Right != nil && !isLeafNode(stack[i].Right){
				stack = append(stack, stack[i].Right)
			}
		}
		stack = stack[n:]
	}
	return sum
}

func isLeafNode(node *TreeNode)bool{
	return node.Right == nil && node.Left == nil
}