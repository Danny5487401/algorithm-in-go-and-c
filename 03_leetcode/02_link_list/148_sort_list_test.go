package _2_link_list

import (
	"github.com/smartystreets/goconvey/convey"
	"sort"
	"testing"
)

// 排序链表 https://leetcode.cn/problems/sort-list/description/?envType=study-plan-v2&envId=top-100-liked

func TestSortList(t *testing.T) {
	convey.Convey("排序链表：按 升序 排列并返回排序后的链表", t, func() {
		testCase := []struct {
			input  *ListNode
			target *ListNode
		}{
			{
				getListNodeBySlice([]int{4, 2, 1, 3}),
				getListNodeBySlice([]int{1, 2, 3, 4}),
			},
			{
				getListNodeBySlice([]int{-1, 5, 3, 4, 0}),
				getListNodeBySlice([]int{-1, 0, 3, 4, 5}),
			},
			{
				getListNodeBySlice([]int{}),
				getListNodeBySlice([]int{}),
			},
		}

		for _, tst := range testCase {
			rsp := sortList(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func sortList(head *ListNode) *ListNode {
	// 要求 O(n log n) 时间复杂度 所以使用归并排序
	return sysSort(head, nil)
}

func sysSort(head, tail *ListNode) *ListNode {

	// 以中点为分界，将链表拆分成两个子链表
	if head == nil {
		return head
	}

	if head.Next == tail {
		head.Next = nil
		return head
	}

	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	mid := slow
	return mergeTwoLists(sysSort(head, mid), sysSort(mid, tail))
}

func sortList1(head *ListNode) *ListNode {
	sliceInfo := make([]int, 0)
	cur := head
	for cur != nil {
		sliceInfo = append(sliceInfo, cur.Val)
		cur = cur.Next
	}
	sort.Ints(sliceInfo)
	var dummy = &ListNode{}
	newOne := dummy
	for i := 0; i < len(sliceInfo); i++ {
		newOne.Next = &ListNode{Val: sliceInfo[i]}
		newOne = newOne.Next
	}
	return dummy.Next
}
