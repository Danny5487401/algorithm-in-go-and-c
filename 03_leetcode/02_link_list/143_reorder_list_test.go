package _2_link_list

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 重排链表 https://leetcode.cn/problems/reorder-list/description/

func TestReorderList(t *testing.T) {
	convey.Convey("重排链表：L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …", t, func() {
		testCase := []struct {
			input  *ListNode
			target *ListNode
		}{
			{
				getListNodeBySlice([]int{1, 2, 3, 4}),
				getListNodeBySlice([]int{1, 4, 2, 3}),
			},
			{
				getListNodeBySlice([]int{1, 2, 3, 4, 5}),
				getListNodeBySlice([]int{1, 5, 2, 4, 3}),
			},
			{
				getListNodeBySlice([]int{1}),
				getListNodeBySlice([]int{1}),
			},
		}

		for _, tst := range testCase {
			reorderList(tst.input)
			convey.So(tst.input, convey.ShouldEqual, tst.target)
		}
	})

}

func reorderList(head *ListNode) {
	// 876.取中间节点
	mid := middleNode(head)

	// 206. 反转
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
