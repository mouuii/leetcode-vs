/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除排序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var i int

	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}

	}
	return i + 1
}

// @lc code=end

