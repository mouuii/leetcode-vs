/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	// 1.暴力法
	// for i := 0; i < a.length-1; i++ {
	// 	for j := i + 1; j < a.length; j++ {

	// 	}
	// }
	// 2.双指针？
	max, start, end := 0, 0, len(height)-1
	for start < end {
		width := end - start
		high := 0

		if height[start] < height[end] {
			height = height[start]
			start++
		} else {
			height = height[left]
			end--
		}
		tmp := width * high
		if temp > max {
			max = temp
		}
		return max
	}
}

// @lc code=end

