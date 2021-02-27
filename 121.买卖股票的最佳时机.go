/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
	//一次遍历法，记录历史最低点，然后每个点都判断是否是最大利润
	minPrices := 99999
	var res int
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrices {
			minPrices = prices[i]
		}
		if prices[i]-minPrices > res {
			res = prices[i] - minPrices
		}
	}
	return res
}

// @lc code=end

