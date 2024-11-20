package _2_link_list

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 相交链表 https://leetcode.cn/problems/intersection-of-two-linked-lists/description/?envType=study-plan-v2&envId=top-100-liked
func TestGetIntersectionNode(t *testing.T) {
	convey.Convey("相交链表", t, func() {
		testCase := []struct {
			input1 *ListNode
			input2 *ListNode
			target *ListNode
		}{
			{
				GetListNodeBySlice([]int{4, 1, 8, 4, 5}),
				GetListNodeBySlice([]int{5, 6, 1}), // 相交 8 4 5
				nil,
			},
			{
				GetListNodeBySlice([]int{1, 9, 1, 2, 4}),
				GetListNodeBySlice([]int{3}), // 相交 2 4
				nil,
			},
		}
		testCase[0].input2.Next.Next.Next = testCase[0].input1.Next.Next
		testCase[0].target = testCase[0].input1.Next.Next

		testCase[0].input2.Next = testCase[0].input1.Next.Next.Next
		testCase[0].target = testCase[0].input2.Next

		for _, tst := range testCase {
			rsp := getIntersectionNode(tst.input1, tst.input2)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

/*
(x+z)+y=(y+z)+X

1. 初始化两个指针 p=headA,q=headB。
2. 不断循环，直到 p=q。
3. 每次循环，p 和 q 各向后走一步。具体来说，如果 p 不是空节点，那么更新 p 为 p.next，否则更新 p 为 headB；如果 q 不是空节点，那么更新 q 为 q.next，否则更新 q 为 headA。
4. 循环结束时，如果两条链表相交，那么此时 p 和 q 都在相交的起始节点处，返回 p；如果两条链表不相交，那么 p 和 q 都走到空节点，所以也可以返回 p，即空节点。

*/

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var p, q = headA, headB
	for p != q {
		if p != nil {
			p = p.Next
		} else {
			p = headB
		}

		if q != nil {
			q = q.Next
		} else {
			q = headA
		}

	}
	return p
}
