package TreeProblems

/*给定一个非空二叉树, 返回一个由每层节点平均值组成的数组。

示例 1：

输入：
3
/ \
9  20
/  \
15   7
输出：[3, 14.5, 11]
解释：
第 0 层的平均值是 3 ,  第1层是 14.5 , 第2层是 11 。因此返回 [3, 14.5, 11] 。*/

func averageOfLevels(root *TreeNode) []float64 {
	res := make([]float64,0)
	if root == nil{
		return res
	}
	stack := [] *TreeNode{root}
	for len(stack) != 0{
		n := len(stack)
		sum, count := 0,0
		for i:=0;i<n;i++{
			sum += stack[i].Val
			count += 1
			if stack[i].Left != nil{
				stack = append(stack, stack[i].Left)
			}
			if stack[i].Right != nil{
				stack = append(stack, stack[i].Right)
			}
		}
		stack = stack[n:]
		res = append(res, float64(sum)/float64(count))
	}
	return res
}