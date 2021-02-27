/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	list := help(root)
	for i := 1; i < len(list); i++ {
		prev, curr := list[i-1], list[i]
		prev.Left, prev.Right = nil, curr
	}

}

func help(root *TreeNode) []*TreeNode {
	list := []*TreeNode{}
	if root != nil {
		list = append(list, root)
		list = append(list, help(root.Left)...)
		list = append(list, help(root.Right)...)

	}
	return list
}

// @lc code=end

