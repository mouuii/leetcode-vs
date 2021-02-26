/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
// time limit
// func climbStairs(n int)int{
// 	if n == 1 || n == 2 {
// 		return n
// 	}
// 	return climbStairs(n-1) + climbStairs(n-2)
// }

// @lc code=end

递归模板
func recursion(level,param...){
	// 1. recursion terminator
	if level > max_level{
		process_result()
		return
	}
	// 2. process logic in current level
	process(level,data...)

	// 3.drill down
	recursion(level+1,param...)
	// 4.reverse the current level stats
}

