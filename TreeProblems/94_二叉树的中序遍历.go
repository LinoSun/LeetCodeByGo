package TreeProblems

/*给定一个二叉树，返回它的中序 遍历。

示例:

输入: [1,null,2,3]
1
\
2
/
3

输出: [1,3,2]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？
*/
func inorderTraversal(root *TreeNode) []int {
	if root == nil{
		return nil
	}
	res := []int{}
	if root.Left != nil{
		res = append(res,inorderTraversal(root.Left)...)
	}
	res = append(res,root.Val)

	if root.Right != nil{
		res = append(res,inorderTraversal(root.Right)...)
	}

	return res
}