package _2_link_list

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 两数相加 II https://leetcode.cn/problems/add-two-numbers-ii/
func TestAddTwoNumbersII(t *testing.T) {
	convey.Convey("两数相加 II:将这两数相加会返回一个新的链表", t, func() {
		testCase := []struct {
			input1 *ListNode
			input2 *ListNode
			target *ListNode
		}{
			{
				GetListNodeBySlice([]int{7, 2, 4, 3}),
				GetListNodeBySlice([]int{5, 6, 4}),
				GetListNodeBySlice([]int{7, 8, 0, 7}),
			},

			{
				GetListNodeBySlice([]int{2, 4, 3}),
				GetListNodeBySlice([]int{5, 6, 4}),
				GetListNodeBySlice([]int{8, 0, 7}),
			},
		}

		for _, tst := range testCase {
			rsp := addTwoNumbersII(tst.input1, tst.input2)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func addTwoNumbersII(l1 *ListNode, l2 *ListNode) *ListNode {
	/*
		反转链表+两数相加
	*/
	l1New := reverseList(l1)
	l2New := reverseList(l2)

	merge := addTwo(l1New, l2New, 0)

	return reverseList(merge)
}
