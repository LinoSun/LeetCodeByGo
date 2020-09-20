package Others
/*给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例:

输入: nums = [1,2,3]
输出:
[
[3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]*/
var res [][]int
func subsets(nums []int) [][]int {
	res = make([][]int,0)
	for i:=0;i<=len(nums);i++{
		cur := make([]int,0,i)
		backtrack77(0,i,cur,nums)
	}
	return res
}

func backtrack77(pos,num int,cur,nums []int){
	//回溯结束
	if len(cur) == num{
		tmp := make([]int,len(cur))
		copy(tmp,cur)
		res = append(res,tmp)
		return
	}
	for i:=pos;i<len(nums);i++{
		cur = append(cur,nums[i])
		backtrack77(i+1,num,cur,nums)
		cur = cur[:len(cur)-1]
	}
}

