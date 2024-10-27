package _2_link_list

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 翻倍以链表形式表示的数字 https://leetcode.cn/problems/double-a-number-represented-as-a-linked-list/description/
func TestDoubleIt(t *testing.T) {
	convey.Convey(" 翻倍以链表形式表示的数字", t, func() {
		testCase := []struct {
			input1 *ListNode
			target *ListNode
		}{
			{
				GetListNodeBySlice([]int{1, 8, 9}),
				GetListNodeBySlice([]int{3, 7, 8}),
			},
			{
				GetListNodeBySlice([]int{9, 9, 9}),
				GetListNodeBySlice([]int{1, 9, 9, 8}),
			},
		}

		for _, tst := range testCase {
			rsp := doubleIt(tst.input1)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func doubleIt(head *ListNode) *ListNode {
	l1New := reverseList(head)
	l2New := addSomeList(l1New, 0)
	return reverseList(l2New)

}

// 递归方式
func addSomeList(l1 *ListNode, carry int) *ListNode {
	if l1 == nil {
		if carry != 0 {
			return &ListNode{Val: carry}
		}
		return nil
	}
	var sum int
	sum = l1.Val*2 + carry
	l1.Val = sum % 10
	l1.Next = addSomeList(l1.Next, sum/10)
	return l1
}
