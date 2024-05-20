package _2_link_list

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 删除排序链表中的重复元素 https://leetcode.cn/problems/remove-duplicates-from-sorted-list/description/
func TestDeleteDuplicates(t *testing.T) {
	convey.Convey("deleteDuplicates ", t, func() {
		testCase := []struct {
			input  *ListNode
			target *ListNode
		}{
			{
				&ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 2,
						},
					},
				},
				&ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
					},
				},
			},
			{
				&ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 3,
								Next: &ListNode{
									Val: 3,
								},
							},
						},
					},
				},
				&ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
						},
					},
				},
			},
		}

		for _, tst := range testCase {
			rsp := deleteDuplicates(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var cur = head

	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next // 删除重复的下一节点
		} else {
			cur = cur.Next // 向右移动
		}
	}
	return head
}
