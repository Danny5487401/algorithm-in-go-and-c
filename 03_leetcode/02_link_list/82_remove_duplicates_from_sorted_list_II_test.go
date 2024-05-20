package _2_link_list

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 删除排序链表中的重复元素II https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/description/
func TestDeleteDuplicatesII(t *testing.T) {
	convey.Convey("deleteDuplicates II", t, func() {
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
								Val: 3,
								Next: &ListNode{
									Val: 4,
									Next: &ListNode{
										Val: 4,
										Next: &ListNode{
											Val: 5,
										},
									},
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
							Val: 5,
						},
					},
				},
			},
			{
				&ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 1,
							Next: &ListNode{
								Val: 2,
								Next: &ListNode{
									Val: 3,
								},
							},
						},
					},
				},
				&ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
		}

		for _, tst := range testCase {
			rsp := deleteDuplicatesII(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func deleteDuplicatesII(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var dummmy = &ListNode{
		Next: head,
	}

	var cur = dummmy

	for cur.Next != nil && cur.Next.Next != nil {
		nextVal := cur.Next.Val
		if nextVal == cur.Next.Next.Val {
			for cur.Next != nil && nextVal == cur.Next.Val {
				// 再下一节点循环判断
				cur.Next = cur.Next.Next
			}

		} else {
			cur = cur.Next
		}
	}

	return dummmy.Next
}
