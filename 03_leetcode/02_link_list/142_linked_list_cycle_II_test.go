package _2_link_list

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 环形链表II https://leetcode.cn/problems/linked-list-cycle-ii/?envType=study-plan-v2&envId=top-100-liked
func TestDetectCycle(t *testing.T) {
	convey.Convey("环形链表", t, func() {
		node1 := &ListNode{Val: 3}
		node2 := &ListNode{Val: 2}
		node1.Next = node2
		node1.Next.Next = &ListNode{Val: 0}
		node1.Next.Next.Next = &ListNode{Val: -4}
		node1.Next.Next.Next.Next = node1.Next

		case2 := &ListNode{Val: 1}
		case2.Next = &ListNode{Val: 2}
		case2.Next.Next = case2

		case3 := &ListNode{Val: 1}
		testCase := []struct {
			input  *ListNode
			target *ListNode
		}{
			{
				node1,
				node2,
			},
			{
				case2,
				case2,
			},
			{
				case3,
				nil,
			},
		}

		for _, tst := range testCase {
			rsp := detectCycle(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}
func detectCycle(head *ListNode) *ListNode {
	var slow, fast = head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow { // 此时快慢相遇
			for slow != head { // 最终会相遇
				// 公式 a-c =(k-1)(b+c)
				// slow 从相遇点出发，head 从头节点出发，两个都走c步，slow 到入口，head 到入口的距离正好是环长的倍数，所以两个都继续走，肯定在入口相遇
				head = head.Next
				slow = slow.Next
			}
			return slow
		}
	}
	return nil
}
