package _2_link_list

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 在链表中插入最大公约数 https://leetcode.cn/problems/insert-greatest-common-divisors-in-linked-list/description/
func TestInsertGreatestCommonDivisors(t *testing.T) {
	convey.Convey("在链表中插入最大公约数: 两个数的 最大公约数 是可以被两个数字整除的最大正整数", t, func() {
		testCase := []struct {
			input  *ListNode
			target *ListNode
		}{
			{
				getListNodeBySlice([]int{18, 6, 10, 3}),
				getListNodeBySlice([]int{18, 6, 6, 2, 10, 1, 3}),
			},
			{
				getListNodeBySlice([]int{7}),
				getListNodeBySlice([]int{7}),
			},
		}

		for _, tst := range testCase {
			rsp := insertGreatestCommonDivisors(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}
func insertGreatestCommonDivisors(head *ListNode) *ListNode {

	// 链表中结点数目在 [1, 5000] 之间。
	var cur = head
	
	for cur.Next != nil {
		newVal := gcd(cur.Val, cur.Next.Val)

		tmp := cur.Next
		cur.Next = &ListNode{Val: newVal}
		cur.Next.Next = tmp

		cur = tmp

	}
	return head
}

// greate-common-divisor
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
