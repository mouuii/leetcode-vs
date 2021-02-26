/*
 * @lc app=leetcode.cn id=589 lang=golang
 *
 * [589] N叉树的前序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
	var ret []int

	var helper func(root *Node)
	helper = func(root *Node) {
		if root == nil {
			return
		}

		ret = append(ret, root.Val)
		for i := range root.Children {
			helper(root.Children[i])
		}
	}

	helper(root)
	return ret
}

// type Stack struct {
// 	d    []*Node
// 	size int
// }

// func (s *Stack) Push(i *Node) {
// 	s.d = append(s.d, i)
// 	s.size++
// }

// func (s *Stack) Pop() (ret *Node) {
// 	s.size = s.size - 1
// 	ret = s.d[s.size]
// 	s.d = s.d[:s.size]
// 	return
// }

// func preorder(root *Node) []int {
// 	if root == nil {
// 		return []int{}
// 	}

// 	stack := Stack{}
// 	var ret []int

// 	stack.Push(root)
// 	for stack.size > 0 {
// 		r := stack.Pop()
// 		ret = append(ret, r.Val)
// 		for i := len(r.Children) - 1; i >= 0; i-- {
// 			stack.Push(r.Children[i])
// 		}
// 	}

// 	return ret
// }

// @lc code=end

