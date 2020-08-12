package stack
/*给定两个 没有重复元素 的数组 nums1 和 nums2 ，其中nums1 是 nums2 的子集。找到 nums1 中每个元素在 nums2 中的下一个比其大的值。

nums1 中数字 x 的下一个更大元素是指 x 在 nums2 中对应位置的右边的第一个比 x 大的元素。如果不存在，对应位置输出 -1 。

 

示例 1:

输入: nums1 = [4,1,2], nums2 = [1,3,4,2].
输出: [-1,3,-1]
解释:
对于num1中的数字4，你无法在第二个数组中找到下一个更大的数字，因此输出 -1。
对于num1中的数字1，第二个数组中数字1右边的下一个较大数字是 3。
对于num1中的数字2，第二个数组中没有下一个更大的数字，因此输出 -1。
*/

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := make([]int,0)
	maps := make(map[int]int)
	for i,v := range nums2{
		if len(stack) == 0 || v < nums2[stack[len(stack)-1]]{
			stack = append(stack, i)
			continue
		}
		for len(stack) != 0 && v>nums2[stack[len(stack)-1]]{
			maps[nums2[stack[len(stack)-1]]] = v
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	res:=make([]int,len(nums1))
	for i,v := range nums1{
		if value,ok := maps[v];ok{
			res[i] = value
		} else {
			res[i] = -1
		}
	}
	return res
}