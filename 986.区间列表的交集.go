/*
 * @lc app=leetcode.cn id=986 lang=golang
 *
 * [986] 区间列表的交集
 */

// @lc code=start
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	res := [][]int{}
	i, j := 0, 0
	for i < len(firstList) && j < len(secondList) {
		start := max(firstList[i][0], secondList[j][0])
		end := min(firstList[i][1], secondList[j][1])
		if start <= end {
			res = append(res, []int{start, end})
		}
		if firstList[i][1] < secondList[j][1] {
			i++
		} else {
			j++
		}
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

