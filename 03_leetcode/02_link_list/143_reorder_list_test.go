package _2_link_list

// 重排链表 https://leetcode.cn/problems/reorder-list/description/

func reorderList(head *ListNode) {
	// 取中间节点
	mid := middleNode(head)

	// 反转
	head2 := reverseList(mid)

	for head2.Next != nil {
		// 记录下个节点
		nxt := head.Next
		nxt2 := head2.Next

		// 开始更改指向
		head.Next = head2
		head2.Next = nxt

		// 更新循环
		head = nxt
		head2 = nxt2
	}

}
