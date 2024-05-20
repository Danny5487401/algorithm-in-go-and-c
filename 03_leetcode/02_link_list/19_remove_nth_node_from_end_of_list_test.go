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
				&ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 5,
								},
							},
						},
					},
				},
				2,
				&ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
			{
				&ListNode{
					Val: 1,
				},
				1,
				nil,
			},
		}

		for _, tst := range testCase {
			rsp := removeNthFromEnd(tst.input, tst.n)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 针对 N 长度
	var dummyNode = &ListNode{
		Next: head,
	}

	// 获取倒数 第 N+1 个结点

	right := dummyNode
	left := dummyNode
	for i := 0; i < n; i++ {
		right = right.Next // 右指针先向右走 n 步,
	}

	for right.Next != nil {
		// 让左右指针的举例始终为 N
		left = left.Next
		right = right.Next
	}
	left.Next = left.Next.Next
	return dummyNode.Next
}
