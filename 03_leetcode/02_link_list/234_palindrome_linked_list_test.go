package _2_link_list

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 回文链表 https://leetcode.cn/problems/palindrome-linked-list/
func TestIsPalindrome(t *testing.T) {
	convey.Convey("重排链表：L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …", t, func() {
		testCase := []struct {
			input  *ListNode
			target bool
		}{
			{
				getListNodeBySlice([]int{1, 2, 2, 1}),
				true,
			},
			{
				getListNodeBySlice([]int{1, 2}),
				false,
			},
		}

		for _, tst := range testCase {
			rsp := isPalindrome(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})
}

func isPalindrome(head *ListNode) bool {
	mid := middleNode(head)
	head2 := reverseList(mid) // 这里分成两段链表
	for head2 != nil {
		if head.Val != head2.Val {
			return false
		}
		head = head.Next
		head2 = head2.Next
	}
	return true

}
