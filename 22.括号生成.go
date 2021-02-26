/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) []string {
	result := []string{}
	generate(0, 2*n, "", result)
}
func generate(level int, max int, s string, result []string) {
	//terminate
	if level >= max {
		fmt.Println(s)
		return
	}
	//process
	s1 := s + "("
	s2 := s + ")"
	//drill down
	generate(level+1, max, s1)
	generate(level+1, max, s2)
	//reverse
}

// @lc code=end

