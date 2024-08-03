package _2_link_list

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 删除链表的倒数第 N 个结点 https://leetcode.cn/problems/remove-nth-node-from-end-of-list/
func TestRemoveNthFromEnd(t *testing.T) {
	convey.Convey("删除链表的倒数第 N 个结点 ", t, func() {
		testCase := []struct {
			input  *ListNode
			n      int
			target *ListNode
		}{
			{
				GetListNodeBySlice([]int{1, 2, 3, 4, 5}),
				2,
				getListNodeBySlice([]int{1, 2, 3, 5}),
			},
			{
				&ListNode{
					Val: 1,
				},
				1,
				getListNodeBySlice([]int{}),
			},
		}

		for _, tst := range testCase {
			rsp := removeNthFromEnd(tst.input, tst.n)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 让左右指针距离始终为N

	// 初始化 dummy node
	var dummyNode = &ListNode{
		Next: head,
	}

	// 初始化指向 dummy node
	right := dummyNode
	left := dummyNode

	for i := 0; i < n; i++ {
		right = right.Next // 右指针先向右走 n 步,
	}

	for right.Next != nil { // 右指针走到最后一个节点
		// 让左右指针的距离始终为 N
		left = left.Next
		right = right.Next
	}
	// 删除节点
	left.Next = left.Next.Next

	return dummyNode.Next
}
