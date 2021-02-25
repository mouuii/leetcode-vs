/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */

// @lc code=start
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	slice := make([]string, numRows)
	index := 0
	sort := true
	numRows--
	for i := 0; i < len(s); i++ {
		slice[index] = slice[index] + string(s[i])
		if index == 0 {
			sort = true
		} else if index == numRows {
			sort = false
		}
		if sort {
			index++
		} else {
			index--
		}
	}
	return strings.Join(slice, "")
}

// @lc code=end

