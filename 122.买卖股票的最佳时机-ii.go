/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start
func maxProfit(prices []int) int {
	var maxProfit int
	for i := 0; i < len(prices); i++ {
		if i != 0 && prices[i]-prices[i-1] > 0 {
			maxProfit += prices[i] - prices[i-1]
		}
	}
	return maxProfit
}

// @lc code=end

