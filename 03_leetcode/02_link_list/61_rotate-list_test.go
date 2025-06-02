package _2_link_list

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 旋转链表 https://leetcode.cn/problems/rotate-list/
func TestRotateRight(t *testing.T) {
	convey.Convey("旋转链表:将链表每个节点向右移动 k 个位置", t, func() {
		testCase := []struct {
			input  *ListNode
			n      int
			target *ListNode
		}{
			{
				GetListNodeBySlice([]int{1, 2, 3, 4, 5}),
				2,
				getListNodeBySlice([]int{4, 5, 1, 2, 3}),
			},
			{
				GetListNodeBySlice([]int{0, 1, 2}),
				4,
				getListNodeBySlice([]int{2, 0, 1}),
			},
			{
				GetListNodeBySlice([]int{1, 2}),
				0,
				getListNodeBySlice([]int{1, 2}),
			},
			{
				GetListNodeBySlice([]int{1, 2}),
				1,
				getListNodeBySlice([]int{2, 1}),
			},
		}

		for _, tst := range testCase {
			rsp := rotateRight(tst.input, tst.n)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

// 方法一：闭合为环
func rotateRight(head *ListNode, k int) *ListNode {
	// 0 <= k <= 2 * 109
	// 链表中节点的数目在范围 [0, 500] 内
	if head == nil || k == 0 || head.Next == nil {
		return head
	}

	// 获取尾节点
	n := 1
	tail := head
	for tail.Next != nil {
		tail = tail.Next
		n++
	}
	// 注意向右移动
	add := n - k%n // 找到新链表的最后一个节点（即原链表的第 (n−1)−(k % n) 个节点）
	if add == n {  // 当链表长度不大于 1，或者 k 为 n 的倍数时，新链表将与原链表相同，我们无需进行任何处理。
		return head
	}

	tail.Next = head

	for i := 1; i <= add; i++ {
		tail = tail.Next // 找到新链表的最后一个节点（即原链表的第 (n−1)−(k % n) 个节点）
	}
	newHead := tail.Next
	tail.Next = nil
	return newHead

}
