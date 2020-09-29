package TreeProblems
/*给定一个二叉树，返回它的 后序 遍历。

示例:

输入: [1,null,2,3]
1
\
2
/
3

输出: [3,2,1]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？
*/
func postorderTraversal(root *TreeNode) []int {
	res := make([]int,0)
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil{
			return
		}
		dfs(root.Left)
		dfs(root.Right)
		res = append(res,root.Val)
	}
	dfs(root)
	return res
}