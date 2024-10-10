package _2_link_list

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 删除排序链表中的重复元素II https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/description/
func TestDeleteDuplicatesII(t *testing.T) {
	convey.Convey("删除排序链表中的重复元素 II：删除所有重复，不保留自身", t, func() {
		testCase := []struct {
			input  *ListNode
			target *ListNode
		}{
			{
				getListNodeBySlice([]int{1, 2, 3, 3, 4, 4, 5}),
				getListNodeBySlice([]int{1, 2, 5}),
			},
			{
				getListNodeBySlice([]int{1, 1, 1, 2, 3}),
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

	var dummy = &ListNode{
		Next: head,
	}

	var cur = dummy

	// 每次循环看下个节点和下下个节点
	for cur.Next != nil && cur.Next.Next != nil {
		nextVal := cur.Next.Val
		if nextVal == cur.Next.Next.Val { // 当前下个 == 当前下下个
			// 可能连续2个以上相同

			for cur.Next != nil && nextVal == cur.Next.Val {
				// 再下一节点循环判断
				cur.Next = cur.Next.Next
			}

		} else {
			cur = cur.Next
		}
	}

	return dummy.Next
}
