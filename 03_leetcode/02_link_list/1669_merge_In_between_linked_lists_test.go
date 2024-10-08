package _2_link_list

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 合并两个链表 https://leetcode.cn/problems/merge-in-between-linked-lists/description/

func TestMergeInBetween(t *testing.T) {
	convey.Convey("合并两个链表:将 list1 中下标从 a 到 b 的全部节点都删除，并将list2 接在被删除节点的位置", t, func() {
		testCase := []struct {
			list1  *ListNode
			a      int
			b      int
			list2  *ListNode
			target *ListNode
		}{
			{
				getListNodeBySlice([]int{10, 1, 13, 6, 9, 5}),
				3, 4,

				getListNodeBySlice([]int{1000000, 1000001, 1000002}),
				getListNodeBySlice([]int{10, 1, 13, 1000000, 1000001, 1000002, 5}),
			},
			{
				getListNodeBySlice([]int{0, 1, 2, 3, 4, 5, 6}),
				2, 5,
				getListNodeBySlice([]int{1000000, 1000001, 1000002, 1000003, 1000004}),
				getListNodeBySlice([]int{0, 1, 1000000, 1000001, 1000002, 1000003, 1000004, 6}),
			},
		}

		for _, tst := range testCase {
			rsp := mergeInBetween(tst.list1, tst.a, tst.b, tst.list2)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	// 1 <= a <= b < list1.length - 1

	var dummy = &ListNode{
		Next: list1,
	}
	left := dummy

	for i := 0; i < a; i++ { //
		left = left.Next
	}

	right := left
	for i := 0; i < b+1-(a-1); i++ { // a 的上一个节点, b 的下一个节点
		right = right.Next
	}
	left.Next = list2
	for list2.Next != nil {
		list2 = list2.Next
	}
	list2.Next = right
	return dummy.Next
}
