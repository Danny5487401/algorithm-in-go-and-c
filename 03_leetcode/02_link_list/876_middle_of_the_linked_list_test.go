package _2_link_list

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 链表的中间节点: https://leetcode.cn/problems/middle-of-the-linked-list/
func TestMiddleNode(t *testing.T) {
	convey.Convey("链表的中间节点:如果是偶数选择第二个节点", t, func() {
		testCase := []struct {
			input  *ListNode
			target *ListNode
		}{
			{
				getListNodeBySlice([]int{1, 2, 3, 4, 5}),
				getListNodeBySlice([]int{3, 4, 5}),
			},
			{
				getListNodeBySlice([]int{1, 2, 3, 4, 5, 6}),
				getListNodeBySlice([]int{4, 5, 6}),
			},
		}

		for _, tst := range testCase {
			rsp := middleNode(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func middleNode(head *ListNode) *ListNode {

	var slow, fast = head, head

	for fast != nil && fast.Next != nil { // fast==nil 偶数，fast.Next==nil 奇数
		// 快指针为空 或则 快指针的下一节点为空 退出循环，上面为反过来判断
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow

}
