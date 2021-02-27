/*
 * @lc app=leetcode.cn id=739 lang=golang
 *
 * [739] 每日温度
 */

// @lc code=start
func dailyTemperatures(T []int) []int {
	length := len(T)
	ans := make([]int, length)
	stack := []int{}
	for i := 0; i < length; i++ {
		temperature := T[i]
		for len(stack) > 0 && temperature > T[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[prevIndex] = i - prevIndex
		}
		stack = append(stack, i)
	}
	return ans
}

// @lc code=end

