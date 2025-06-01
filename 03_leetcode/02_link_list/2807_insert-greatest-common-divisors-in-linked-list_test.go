package _2_link_list

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 在链表中插入最大公约数 https://leetcode.cn/problems/insert-greatest-common-divisors-in-linked-list/description/
func TestInsertGreatestCommonDivisors(t *testing.T) {
	convey.Convey("在链表中插入最大公约数: 在相邻结点之间，请你插入一个新的结点，结点值为这两个相邻结点值的 最大公约数 。"+
		"两个数的 最大公约数 是可以被两个数字整除的最大正整数", t, func() {
		testCase := []struct {
			input  *ListNode
			target *ListNode
		}{
			{
				/*
					- 18 和 6 的最大公约数为 6 ，插入第一和第二个结点之间。
					- 6 和 10 的最大公约数为 2 ，插入第二和第三个结点之间。
					- 10 和 3 的最大公约数为 1 ，插入第三和第四个结点之间
				*/
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
