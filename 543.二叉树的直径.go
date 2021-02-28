/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
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
var max int

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max = 0
	depth(root)
	return max
}

func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := depth(node.Left)
	right := depth(node.Right)
	max = int(math.Max(float64(left+right), float64(max)))  // 更新全局变量max值
	return int(math.Max(float64(left), float64(right))) + 1 // 返回node节点的深度
}

// @lc code=end

