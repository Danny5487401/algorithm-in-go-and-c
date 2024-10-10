package _2_link_list

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 合并两个有序链表 https://leetcode.cn/problems/merge-two-sorted-lists/description/?envType=study-plan-v2&envId=top-100-liked
func TestMergeTwoLists(t *testing.T) {
	convey.Convey("合并两个有序链表 ", t, func() {
		testCase := []struct {
			list1  *ListNode
			list2  *ListNode
			target *ListNode
		}{
			{

				GetListNodeBySlice([]int{1, 2, 4}),
				GetListNodeBySlice([]int{1, 3, 4}),
				GetListNodeBySlice([]int{1, 1, 2, 3, 4, 4}),
			},
			{
				GetListNodeBySlice([]int{}),
				GetListNodeBySlice([]int{}),
				GetListNodeBySlice([]int{}),
			},
			{
				GetListNodeBySlice([]int{}),
				GetListNodeBySlice([]int{0}),
				GetListNodeBySlice([]int{0}),
			},
		}

		for _, tst := range testCase {
			rsp := mergeTwoLists(tst.list1, tst.list2)
			convey.So(rsp, convey.ShouldResemble, tst.target)
		}
	})

}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var dummyNode = &ListNode{} // 用哨兵节点简化代码逻辑
	cur := dummyNode
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
	// 拼接剩余链表
	if list1 != nil {
		cur.Next = list1
	} else {
		cur.Next = list2
	}
	return dummyNode.Next
}
