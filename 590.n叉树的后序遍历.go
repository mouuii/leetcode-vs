/*
 * @lc app=leetcode.cn id=590 lang=golang
 *
 * [590] N叉树的后序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
	if root != nil {
		rest := make([]int, 0, 1)
		return append(postOrderFunc(root, rest), root.Val)
	} else {
		return nil
	}
}

func postOrderFunc(root *Node, rest []int) []int {
	// 如果没有子树，则返回根节点的值
	if root == nil {
		return []int{}
	}

	// 如果有子树则从右至左查找
	for idx, _ := range root.Children {
		rest = append(postOrderFunc(root.Children[idx], rest), root.Children[idx].Val)
	}
	return rest
}

// @lc code=end

