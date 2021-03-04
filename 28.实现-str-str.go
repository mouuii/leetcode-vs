/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	if haystack == "" && needle == "" {
		return 0
	}
	var index = -1
	r := []rune(haystack)
	length := utf8.RuneCountInString(needle)
	r1 := len(r)
	for i := 0; i < r1; i++ {
		end := i + length
		if end > r1 {
			return index
		}
		sub := append([]rune{}, r[i:end]...)
		if string(sub) == needle {
			return i
		}
	}
	return index
}

// @lc code=end

