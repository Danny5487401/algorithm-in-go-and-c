package _2_link_list

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 环形链表 https://leetcode.cn/problems/linked-list-cycle/description/?envType=study-plan-v2&envId=top-100-liked
func TestHasCycle(t *testing.T) {
	convey.Convey("环形链表", t, func() {
		exampleNode1 := &ListNode{Val: 3}
		exampleNode1.Next = &ListNode{Val: 2}
		exampleNode1.Next.Next = &ListNode{Val: 0}
		exampleNode1.Next.Next.Next = &ListNode{Val: -4}
		exampleNode1.Next.Next.Next.Next = exampleNode1.Next

		exampleNode2 := &ListNode{Val: 1}
		exampleNode2.Next = &ListNode{Val: 2}
		exampleNode2.Next.Next = exampleNode2

		exampleNode3 := &ListNode{Val: 1}
		testCase := []struct {
			input  *ListNode
			target bool
		}{
			{
				exampleNode1,
				true,
			},
			{
				exampleNode2,
				true,
			},
			{
				exampleNode3,
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
