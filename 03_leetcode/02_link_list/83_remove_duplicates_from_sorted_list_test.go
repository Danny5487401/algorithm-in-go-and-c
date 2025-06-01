package _2_link_list

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 删除排序链表中的重复元素 https://leetcode.cn/problems/remove-duplicates-from-sorted-list/description/
func TestDeleteDuplicates(t *testing.T) {
	convey.Convey("删除排序链表中的重复元素:已排序的链表的头 head ， 删除所有重复的元素，使每个元素只出现一次 。保留一个重复 ", t, func() {
		testCase := []struct {
			input  *ListNode
			target *ListNode
		}{
			{
				getListNodeBySlice([]int{1, 1, 2}),
				getListNodeBySlice([]int{1, 2}),
			},
			{
				getListNodeBySlice([]int{1, 1, 2, 3, 3}),
				getListNodeBySlice([]int{1, 2, 3}),
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

	// 可以保留头节点，所以可以不用 dummy node

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
