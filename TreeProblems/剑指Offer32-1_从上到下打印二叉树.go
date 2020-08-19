package TreeProblems
/*从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。

 

例如:
给定二叉树: [3,9,20,null,null,15,7],

3
/ \
9  20
/  \
15   7
返回：

[3,9,20,15,7]*/

func levelOrder(root *TreeNode) []int {
	var res []int
	if root==nil{
		return res
	}
	queue := []*TreeNode{root}
	for len(queue)>0{
		size :=len(queue)
		for i:=0;i<size;i++{
			node :=queue[i]
			if node.Left!=nil{
				queue = append(queue,node.Left)
			}
			if node.Right!=nil{
				queue = append(queue,node.Right)
			}
			res = append(res,node.Val)
		}
		queue = queue[size:]
	}
	return res
}