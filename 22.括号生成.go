/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) []string {
	result := []string{}
	generate(0, 0, n, "", &result)
	return result
}

func generate(left int, right int, n int, s string, result *[]string) {
	//terminate
	if left == n && right == n {
		*result = append(*result, s)
		return
	}
	//process
	s1 := s + "("
	s2 := s + ")"
	//drill down
	if left < n {
		generate(left+1, right, n, s1, result)
	}
	if left > right {
		generate(left, right+1, n, s2, result)
	}

	//reverse
}

// @lc code=end

