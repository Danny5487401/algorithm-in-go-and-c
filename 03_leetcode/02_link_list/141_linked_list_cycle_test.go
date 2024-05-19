package _2_link_list

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 环形链表 https://leetcode.cn/problems/linked-list-cycle/description/?envType=study-plan-v2&envId=top-100-liked
func TestHasCycle(t *testing.T) {
	convey.Convey("环形链表", t, func() {
		node1 := &ListNode{Val: 3}
		node1.Next = &ListNode{Val: 2}
		node1.Next.Next = &ListNode{Val: 0}
		node1.Next.Next.Next = &ListNode{Val: -4}
		node1.Next.Next.Next.Next = node1.Next

		node2 := &ListNode{Val: 1}
		node2.Next = &ListNode{Val: 2}
		node2.Next.Next = node2

		node3 := &ListNode{Val: 1}
		testCase := []struct {
			input  *ListNode
			target bool
		}{
			{
				node1,
				true,
			},
			{
				node2,
				true,
			},
			{
				node3,
				false,
			},
		}

		for _, tst := range testCase {
			rsp := hasCycle(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func hasCycle(head *ListNode) bool {
	var slow, fast = head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			// 相遇
			return true
		}
	}
	return false

}
