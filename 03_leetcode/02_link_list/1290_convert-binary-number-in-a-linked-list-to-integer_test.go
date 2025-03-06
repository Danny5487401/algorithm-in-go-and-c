package _2_link_list

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 二进制链表转整数 https://leetcode.cn/problems/convert-binary-number-in-a-linked-list-to-integer/description/
func TestGetDecimalValue(t *testing.T) {
	convey.Convey(" 二进制链表转整数", t, func() {
		testCase := []struct {
			input  *ListNode
			target int
		}{
			{
				getListNodeBySlice([]int{1, 0, 1}),
				5,
			},
			{
				getListNodeBySlice([]int{0}),
				0,
			},
			{
				getListNodeBySlice([]int{1}),
				1,
			},
			{
				getListNodeBySlice([]int{1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0}),
				18880,
			},
		}

		for _, tst := range testCase {
			rsp := getDecimalValue(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}
func getDecimalValue(head *ListNode) int {
	if head == nil {
		return 0
	}
	var ans int
	for head != nil {
		ans = ans*2 + head.Val
		head = head.Next
	}
	return ans

}
