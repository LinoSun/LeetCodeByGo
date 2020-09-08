package Others

/*给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

示例:

输入: n = 4, k = 2
输出:
[
[2,4],
[3,4],
[2,3],
[1,2],
[1,3],
[1,4],
]*/


func combine(n int, k int) [][]int {
	var ans [][]int

	var dfs func(int, []int)
	dfs = func(m int, path []int) {
		if len(path) == k {
			tmp := make([]int, k)
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}
		if m > n {
			return
		}

		dfs(m + 1, append(path, m))
		dfs(m + 1, path)
		return
	}

	dfs(1, []int{})
	return ans
}
