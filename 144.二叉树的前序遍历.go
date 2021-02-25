/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
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
func preorderTraversal(root *TreeNode) []int {
	var result []int
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		result = append(result, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return result
}

// 试试非递归吧这里
// func preorderTraversal2(root *TreeNode) []int {
// 	if root == nil {
// 		return []int{}
// 	}
// 	stack, res := []*TreeNode{}, []int{}
// 	stack = append(stack, root)
// 	for len(stack) != 0 {
// 		node := stack[len(stack)-1]
// 		stack = stack[:len(stack)-1]
// 		if node != nil {
// 			res = append(res, node.Val)
// 		}
// 		if node.Right != nil {
// 			stack = append(stack, node.Right)
// 		}
// 		if node.Left != nil {
// 			stack = append(stack, node.Left)
// 		}
// 	}
// 	return res
// }

// @lc code=end

