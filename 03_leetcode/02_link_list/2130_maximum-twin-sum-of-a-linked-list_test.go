package _2_link_list

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 链表最大孪生和 https://leetcode.cn/problems/maximum-twin-sum-of-a-linked-list/description/
func TestPairSum(t *testing.T) {
	convey.Convey("链表最大孪生和", t, func() {
		testCase := []struct {
			input  *ListNode
			target int
		}{
			{
				getListNodeBySlice([]int{5, 4, 2, 1}),
				6,
			},
			{
				getListNodeBySlice([]int{4, 2, 2, 3}),
				7,
			},
		}

		for _, tst := range testCase {
			rsp := pairSum(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func pairSum(head *ListNode) int {
	mid := middleNode(head)
	head2 := reverseList(mid)

	maxValue := 0
	for head2 != nil {
		value := head.Val + head2.Val
		maxValue = max(maxValue, value)
		head2 = head2.Next
		head = head.Next
	}
	return maxValue
}
