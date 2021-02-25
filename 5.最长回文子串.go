/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
// 动态规划
func longestPalindrome(s string) string {
	// 1。找到状态 - 》子串是否是回文
	// 2. 找到转移方程 -》dp[i][j] = (s[i] == s[j]) and dp[i + 1][j - 1]
	// 3. 初始化  -》 初始化的时候，单个字符一定是回文串，因此把对角线先初始化为 true，即 dp[i][i] = true 。

	// 废话不多说，先定一个二维数组来记录状态
	n := len(s)
	ans := ""
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for l := 0; l < n; l++ {
		for i := 0; i+l < n; i++ {
			j := i + l
			if l == 0 {
				dp[i][j] = 1
			} else if l == 1 {
				if s[i] == s[j] {
					dp[i][j] = 1
				}
			} else {
				if s[i] == s[j] {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] > 0 && l+1 > len(ans) {
				ans = s[i : i+l+1]
			}
		}
	}
	return ans

}

// @lc code=end

