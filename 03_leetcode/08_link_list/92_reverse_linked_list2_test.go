package _8_link_list

// 反转链表II https://leetcode.cn/problems/reverse-linked-list-ii/description/

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// 哨兵节点
	var dummy = &ListNode{
		Next: head,
	}
	var p0 = dummy
	for i := 0; i < left-1; i++ {
		p0 = p0.Next
	}

	// 同 反转链表 https://leetcode.cn/problems/reverse-linked-list/?envType=study-plan-v2&envId=top-100-liked
	var cur = p0.Next
	var pre *ListNode = nil
	for i := 0; i < right-left+1; i++ {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt

	}

	// 反转边界
	p0.Next.Next = cur
	p0.Next = pre

	return dummy.Next
}
