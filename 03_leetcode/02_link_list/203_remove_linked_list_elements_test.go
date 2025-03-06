package _2_link_list

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 移除链表元素  https://leetcode.cn/problems/remove-linked-list-elements/
func TestRemoveElements(t *testing.T) {
	convey.Convey("移除链表元素", t, func() {
		testCase := []struct {
			input  *ListNode
			val    int
			target *ListNode
		}{
			{
				getListNodeBySlice([]int{1, 2, 6, 3, 4, 5, 6}),
				6,
				getListNodeBySlice([]int{1, 2, 3, 4, 5}),
			},
			{
				getListNodeBySlice([]int{}),
				1,
				getListNodeBySlice([]int{}),
			},
			{
				getListNodeBySlice([]int{7, 7, 7, 7}),
				7,
				getListNodeBySlice([]int{}),
			},
		}

		for _, tst := range testCase {
			rsp := removeElements(tst.input, tst.val)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	var dummy = &ListNode{Next: head} // 可能删除的第一个节点是头节点

	var cur = dummy

	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next // 删除下一个节点
		} else {
			cur = cur.Next // 继续向后遍历链表
		}

	}
	return dummy.Next
}
