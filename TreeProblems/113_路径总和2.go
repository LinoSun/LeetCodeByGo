package TreeProblems
/*给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。

说明: 叶子节点是指没有子节点的节点。

示例:
给定如下二叉树，以及目标和 sum = 22，

5
/ \
4   8
/   / \
11  13  4
/  \    / \
7    2  5   1
返回:

[
[5,4,11,2],
[5,8,4,5]
]*/
func pathSum(root *TreeNode, sum int) [][]int {
	res := make([][]int,0)
	path := make([]int,0)
	dfs113(&res,root,path,sum)
	return res
}

func dfs113(res *[][]int,root *TreeNode,path []int,left int){
	if root == nil{
		return
	}
	if root.Left == nil && root.Right==nil && root.Val == left{
		tmp := make([]int,len(path)+1)
		copy(tmp,append(path,root.Val))
		*res = append(*res,tmp)
		return
	}
	path = append(path,root.Val)
	dfs113(res,root.Left,path,left-root.Val)
	dfs113(res,root.Right,path,left-root.Val)
}