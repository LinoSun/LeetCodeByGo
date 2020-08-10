package Queue

/*给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。

示例:

输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
输出: [3,3,5,5,6,7]
解释:

滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
1 [3  -1  -3] 5  3  6  7       3
1  3 [-1  -3  5] 3  6  7       5
1  3  -1 [-3  5  3] 6  7       5
1  3  -1  -3 [5  3  6] 7       6
1  3  -1  -3  5 [3  6  7]      7
*/

func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int,0)
	n := len(nums)
	if n == 0{
		return res
	}
	for i:=0;i+k<=n;i++{
		res = append(res, intMax(nums[i:i+k]))
	}
	return res
}

func intMax(nums []int)int{
	res := nums[0]
	for _,v := range nums{
		if v > res{
			res = v
		}
	}
	return res
}