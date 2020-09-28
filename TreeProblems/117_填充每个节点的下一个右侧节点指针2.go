package TreeProblems

//给定一个二叉树

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

/*填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。

初始状态下，所有 next 指针都被设置为 NULL。

 

进阶：

你只能使用常量级额外空间。
使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。
 

示例：



输入：root = [1,2,3,4,5,null,7]
输出：[1,#,2,3,#,4,5,7,#]
解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。
 

提示：

树中的节点数小于 6000
-100 <= node.val <= 100*/

func connect(root *Node) *Node {
	if root == nil{
		return root
	}
	queue := []*Node{root}
	for len(queue) != 0{
		n := len(queue)
		for i:=0;i<n;i++{
			if i+1 == n{
				queue[i].Next = nil
			} else{
				queue[i].Next = queue[i+1]
			}
			if queue[i].Left != nil{
				queue = append(queue,queue[i].Left)
			}
			if queue[i].Right != nil{
				queue = append(queue,queue[i].Right)
			}
		}
		queue = queue[n:]
	}
	return root
}
