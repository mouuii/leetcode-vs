/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	option := make([]int, len(nums))
	visted := make([]int, len(nums))
	var dfs func(int)
	dfs = func(idx int) {
		if idx == len(nums) {
			ans = append(ans, append([]int{}, option...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if visted[i] == 0 {
				option[idx] = nums[i]
				visted[i] = 1
				dfs(idx + 1)
				visted[i] = 0
			}
		}
	}
	dfs(0)
	return ans
}

// @lc code=end

