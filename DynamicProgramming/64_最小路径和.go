package DynamicProgramming

/*给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。

示例:

输入:
[
  [1,3,1],
[1,5,1],
[4,2,1]
]
输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。*/

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rows, columns := len(grid), len(grid[0])
	dp := make([][]int,rows)
	for i := 0;i<rows;i++{
		dp[i] = make([]int,columns)
	}
	dp[0][0] = grid[0][0]
	for i := 1;i<rows;i++{
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for i := 1;i<columns;i++{
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}
	for i := 1;i<rows;i++{
		for j:=1;j<columns;j++{
			dp[i][j] = minPathSumMin(dp[i-1][j] + grid[i][j],dp[i][j-1] + grid[i][j])
		}
	}
	return dp[rows-1][columns-1]
}

func minPathSumMin(x,y int)int{
	if x < y {
		return x
	}
	return y
}
