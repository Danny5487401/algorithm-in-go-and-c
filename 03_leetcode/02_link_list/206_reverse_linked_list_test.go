package _2_link_list

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

// 反转链表 https://leetcode.cn/problems/reverse-linked-list/?envType=study-plan-v2&envId=top-100-liked
func TestReverseList(t *testing.T) {
	convey.Convey("反转链表:将当前节点的下一个节点指向上一个节点", t, func() {
		testCase := []struct {
			input  *ListNode
			target *ListNode
		}{
			{
				GetListNodeBySlice([]int{1, 2, 3, 4, 5}),
				GetListNodeBySlice([]int{5, 4, 3, 2, 1}),
			},

			{
				&ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
					},
				},
				&ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 1,
					},
				},
			},
		}

		for _, tst := range testCase {
			rsp := reverseList(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode = nil

	var cur = head
	for cur != nil {
		// 记录当前节点的下一个节点
		nxt := cur.Next

		// 反转： 将当前节点的下个节点指向上一个节点
		cur.Next = pre

		// 更新 pre 和 cur: 注意顺序
		pre = cur
		cur = nxt
	}

	// 循环结束: cur 指定空,  pre 指定末尾,返回 pre
	return pre

}
