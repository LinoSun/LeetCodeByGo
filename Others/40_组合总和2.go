package Others

import "sort"

/*给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用一次。

说明：

所有数字（包括目标数）都是正整数。
解集不能包含重复的组合。 
示例 1:

输入: candidates = [10,1,2,7,6,1,5], target = 8,
所求解集为:
[
[1, 7],
[1, 2, 5],
[2, 6],
[1, 1, 6]
]
示例 2:

输入: candidates = [2,5,2,1,2], target = 5,
所求解集为:
[
  [1,2,2],
  [5]
]*/

var ans [][]int
var list []int

func combinationSum2(candidates []int, target int) [][]int {
	ans = make([][]int,0)
	list = make([]int,0)
	if len(candidates) == 0{
		return ans
	}
	sort.Ints(candidates)

	combinationSum2Dfs(candidates,0,target)
	return ans
}

func combinationSum2Dfs (candidates []int, start int, target int){
	if target == 0{
		tmp := make([]int,len(list))
		copy(tmp,list)
		ans = append(ans, tmp)
		return
	}

	for i:=start;i<len(candidates);i++{
		// 剪枝 已排序的数组后面的数字差肯定都 <0
		if target - candidates[i] < 0{
			break
		}
		if i > start && candidates[i] == candidates[i-1] {
			continue // 去重，去掉第一个
		}
		list = append(list, candidates[i])
		combinationSum2Dfs(candidates,i+1,target-candidates[i])
		list = list[:len(list)-1]
	}
}