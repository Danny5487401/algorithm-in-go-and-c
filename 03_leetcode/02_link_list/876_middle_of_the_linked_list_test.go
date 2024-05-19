package _2_link_list

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 链表的中间节点: https://leetcode.cn/problems/middle-of-the-linked-list/
func TestMiddleNode(t *testing.T) {
	convey.Convey("链表的中间节点", t, func() {
		testCase := []struct {
			input  *ListNode
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
				&ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
						},
					},
				},
			},
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
									Next: &ListNode{
										Val: 6,
									},
								},
							},
						},
					},
				},
				&ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 6,
						},
					},
				},
			},
		}

		for _, tst := range testCase {
			rsp := middleNode(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}
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
