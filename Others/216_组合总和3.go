package Others

/*找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。

说明：

所有数字都是正整数。
解集不能包含重复的组合。
示例 1:

输入: k = 3, n = 7
输出: [[1,2,4]]
示例 2:

输入: k = 3, n = 9
输出: [[1,2,6], [1,3,5], [2,3,4]]*/

var comSumRes [][]int // 结果集
var comSumList []int // 临时结果集

func combinationSum3(k int, n int) [][]int {
	/*
	   回溯 + 剪枝 + 去重
	   剪枝条件： > k 去掉
	   去重条件： 加 start 控制
	*/
	comSumRes = make([][]int, 0)
	comSumList = make([]int, 0)
	backtarck(k, n, 1)
	return comSumRes
}

func backtarck(k int, n int, start int) {
	if n == 0 && k == 0 {
		tmp := make([]int, len(comSumList))
		copy(tmp, comSumList)
		comSumRes = append(comSumRes, tmp)
		return
	}

	for i:=start;i<=9;i++ {
		if n - i < 0 {
			break
		}
		if k == 1 && n-i > 0 {
			continue
		}
		comSumList = append(comSumList, i)
		backtarck(k-1, n-i, i+1)
		comSumList = comSumList[:len(comSumList)-1]
	}
}