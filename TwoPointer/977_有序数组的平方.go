package TwoPointer
/*给定一个按非递减顺序排序的整数数组 A，返回每个数字的平方组成的新数组，要求也按非递减顺序排序。

 

示例 1：

输入：[-4,-1,0,3,10]
输出：[0,1,9,16,100]
示例 2：

输入：[-7,-3,2,3,11]
输出：[4,9,9,49,121]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/squares-of-a-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func sortedSquares(A []int) []int {
	n := len(A)
	res := make([]int,n)
	l,r,index := 0,n-1,n-1
	for l <= r{
		if A[l] * A[l] > A[r] * A[r]{
			res[index] = A[l] * A[l]
			l++
		} else {
			res[index] = A[r] * A[r]
			r--
		}
		index --
	}
	return res
}