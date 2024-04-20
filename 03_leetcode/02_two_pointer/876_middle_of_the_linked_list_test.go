package _2_two_pinter

// 链表的中间节点: https://leetcode.cn/problems/middle-of-the-linked-list/

func middleNode(head *ListNode) *ListNode {

	var slow = head
	var fast = head

	for fast != nil && fast.Next != nil {
		// 快指针为空 或则 快指针的下一节点为空 退出循环，上面为反过来判断
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow

}

type ListNode struct {
	Val  int
	Next *ListNode
}
