package TreeProblems

/*给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

例如：
给定二叉树 [3,9,20,null,null,15,7],

3
/ \
9  20
/  \
15   7
返回其自底向上的层次遍历为：

[
[15,7],
[9,20],
[3]
]
*/

func levelOrderBottom(root *TreeNode) [][]int {
	queue := []*TreeNode{root}
	res := make([][]int,0)
	if root == nil{
		return res
	}
	for len(queue) != 0{
		n := len(queue)
		temp := make([]int,0)
		for i:=0;i<n;i++{
			temp = append(temp, queue[i].Val)
			if queue[i].Left != nil{
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil{
				queue = append(queue, queue[i].Right)
			}
		}
		res = append([][]int{temp},res...)
		queue = queue[n:]
	}
	return res
}