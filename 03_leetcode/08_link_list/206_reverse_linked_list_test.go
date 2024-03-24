package _8_link_list

// 反转链表 https://leetcode.cn/problems/reverse-linked-list/?envType=study-plan-v2&envId=top-100-liked

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode = nil
	var cur = head
	for cur != nil {
		nxt := cur.Next
		// 反转
		cur.Next = pre
		// 更新 pre 和 cur: 注意顺序
		pre = cur
		cur = nxt
	}

	return pre

}
