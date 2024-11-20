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

// 方式一: 自顶向下归并排序
func sysSort(head, tail *ListNode) *ListNode {

	// 如果链表为空或者只有一个节点，无需排序
	if head == nil || head.Next == nil {
		return head
	}

	mid := cutMiddleNode(head)

	// 分治
	head1 := sysSort(head, mid)
	head2 := sysSort(mid, tail)

	// 合并
	return mergeTwoLists(head1, head2)
}

// 876. 链表的中间结点（快慢指针）
func cutMiddleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	// 先找到链表的中间结点的前一个节点
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next // 下一个节点就是链表的中间结点 mid
	slow.Next = nil  // 断开 mid 的前一个节点和 mid 的连接
	return mid
}

// 方法二：归并排序（迭代）

// 废弃
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
