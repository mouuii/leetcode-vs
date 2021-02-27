package main

func main() {

}

/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存机制
 */

// @lc code=start

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

		return list
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
