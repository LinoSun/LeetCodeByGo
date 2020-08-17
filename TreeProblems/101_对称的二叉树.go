package TreeProblems

/*请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3
但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3
*/

func isSymmetric(root *TreeNode) bool {
	return check(root,root)
}

func check(left,right *TreeNode)bool{
	if left == nil && right == nil{
		return true
	}
	if left == nil || right == nil{
		return false
	}
	return left.Val == right.Val && check(left.Left,right.Right) && check(left.Right,right.Left)
}