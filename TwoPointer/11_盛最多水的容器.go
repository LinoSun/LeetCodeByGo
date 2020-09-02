package TwoPointer

/*给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且 n 的值至少为 2。*/



func maxArea(height []int) int {
	l,r := 0, len(height) - 1
	var area int

	for l < r {
		curArea := min(height[l],height[r]) * (r-l)
		area = max(area,curArea)
		if height[l] < height[r] {
			l += 1
		} else {
			r -= 1
		}
	}
	return area
}

func max(x,y int) int{
	if x > y{
		return x
	}
	return y
}

func min(x,y int)int{
	if x < y{
		return x
	}
	return y
}
