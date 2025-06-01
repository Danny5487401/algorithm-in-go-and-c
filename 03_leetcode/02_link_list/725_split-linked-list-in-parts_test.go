package _2_link_list

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 分隔链表 https://leetcode.cn/problems/split-linked-list-in-parts/description/
func TestSplitListToParts(t *testing.T) {
	convey.Convey("分隔链表:每部分的长度应该尽可能的相等：任意两部分的长度差距不能超过 1 。这可能会导致有些部分为 null 。"+
		"这 k 个部分应该按照在链表中出现的顺序排列，并且排在前面的部分的长度应该大于或等于排在后面的长度", t, func() {
		testCase := []struct {
			input1 *ListNode
			k      int
			target []*ListNode
		}{
			{
				GetListNodeBySlice([]int{1, 2, 3}),
				5,
				[]*ListNode{
					GetListNodeBySlice([]int{1}),
					GetListNodeBySlice([]int{2}),
					GetListNodeBySlice([]int{3}),
					GetListNodeBySlice([]int{}),
					GetListNodeBySlice([]int{}),
				},
			},

			{
				GetListNodeBySlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
				3,
				[]*ListNode{
					GetListNodeBySlice([]int{1, 2, 3, 4}),
					GetListNodeBySlice([]int{5, 6, 7}),
					GetListNodeBySlice([]int{8, 9, 10}),
				},
			},
		}

		for _, tst := range testCase {
			rsp := splitListToParts(tst.input1, tst.k)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func splitListToParts(head *ListNode, k int) []*ListNode {
	n := 0
	for node := head; node != nil; node = node.Next {
		n++
	}
	/*
		得到链表的长度 n 之后，记 quotient(商) =⌊n/k⌋，remainder(余数)=n%k，则在分隔成的 k 个部分中，前 remainder 个部分的长度各为 quotient+1，其余每个部分的长度各为 quotient
	*/
	quotient, remainder := n/k, n%k

	parts := make([]*ListNode, k)
	for i, curr := 0, head; i < k && curr != nil; i++ {
		parts[i] = curr      // 将 curr 作为当前部分的头结点；
		partSize := quotient //计算当前部分的长度 partSize；
		if i < remainder {
			partSize++ // 前 remainder 个部分的长度各为 quotient+1
		}
		for j := 1; j < partSize; j++ {
			curr = curr.Next
		}
		curr, curr.Next = curr.Next, nil // 需要拆分 curr 和后面一个结点之间的连接关系
	}
	return parts
}
