package TreeProblems

import "strconv"

/*给定一个二叉树，返回所有从根节点到叶子节点的路径。

说明: 叶子节点是指没有子节点的节点。

示例:

输入:

1
/   \
2     3
\
5

输出: ["1->2->5", "1->3"]

解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3
*/
var paths []string

func binaryTreePaths(root *TreeNode) []string {
	paths = []string{}
	binaryTreePathsDFS(root, "")
	return paths
}

func binaryTreePathsDFS(root *TreeNode, path string) {
	if root != nil {
		pathSB := path
		pathSB += strconv.Itoa(root.Val)
		if root.Left == nil && root.Right == nil {
			paths = append(paths, pathSB)
		} else {
			pathSB += "->"
			binaryTreePathsDFS(root.Left,pathSB)
			binaryTreePathsDFS(root.Right,pathSB)
		}
	}
}
