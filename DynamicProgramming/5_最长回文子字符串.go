package DynamicProgramming

func longestPalindrome(s string) string {
	n := len(s)
	// 创建二维dp数组
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	ans := ""
	// 枚举出所有的字符串，显示长度
	for l := 0; l < n; l++ {
		// i是起始位置，j是终止位置
		for i := 0; i+l < n; i++ {
			j := i + l
			if l == 0 {
				dp[i][j] = 1
			} else if l == 1 {
				if s[i] == s[j] {
					dp[i][j] = 1 // 边界判断完成
				}
			} else {
				if s[i] == s[j] {
					dp[i][j] = dp[i+1][j-1] // 如果相等就转义状态，不用管上一个状态的值
				}
			}
			if dp[i][j] > 0 && l+1 > len(ans) {
				ans = s[i : i+l+1]
			}
		}
	}
	return ans
}
