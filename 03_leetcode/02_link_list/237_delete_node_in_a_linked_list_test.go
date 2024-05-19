package _2_link_list

// 删除链表中的节点 https://leetcode.cn/problems/delete-node-in-a-linked-list/
func deleteNode(node *ListNode) {
	// 给的节点是要被删除

	// copy 下个节点的值，删除下一个节点
	node.Val = node.Next.Val
	node.Next = node.Next.Next

}
