package TreeProblems

/*给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
示例：
二叉树：[3,9,20,null,null,15,7],

3
/ \
9  20
/  \
15   7
返回其层次遍历结果：

[
[3],
[9,20],
[15,7]
]*/

func levelOrder2(root *TreeNode) [][]int {
	res := make([][]int,0)
	if root == nil{
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) != 0{
		n := len(queue)
		tmp := make([]int,n)
		for i,v := range queue{
			tmp[i] = v.Val
		}
		res = append(res, tmp)
		for i:=0;i<n;i++{
			if queue[i].Left != nil{
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil{
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[n:]
	}
	return res
}

