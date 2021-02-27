/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	result := lists[0]
	for i := 1; i < len(lists); i++ {
		result = Merge(result, lists[i]) //每个都合并
	}
	return result
}

func Merge(l1 *ListNode, l2 *ListNode) *ListNode { //合并链表算法
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var res *ListNode
	if l1.Val >= l2.Val {
		res = l2
		res.Next = Merge(l1, l2.Next)
	} else {
		res = l1
		res.Next = Merge(l1.Next, l2)
	}
	return res
}

// @lc code=end

