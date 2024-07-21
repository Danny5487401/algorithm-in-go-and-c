package _2_link_list

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 反转链表II https://leetcode.cn/problems/reverse-linked-list-ii/description/
func TestReverseBetween(t *testing.T) {
	convey.Convey("反转中间部分链表", t, func() {
		testCase := []struct {
			input       *ListNode
			left, right int
			target      *ListNode
		}{
			{
				GetListNodeBySlice([]int{1, 2, 3, 4, 5}),
				2, 4,
				GetListNodeBySlice([]int{1, 4, 3, 2, 5}),
			},

			{
				&ListNode{
					Val: 5,
				},
				1, 1,
				&ListNode{
					Val: 5,
				},
			},
		}

		for _, tst := range testCase {
			rsp := reverseBetween(tst.input, tst.left, tst.right)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// 哨兵节点： 考虑 left=1 的情况
	var dummy = &ListNode{
		Next: head,
	}

	var p0 = dummy                // p0 反转段的上一个节点
	for i := 0; i < left-1; i++ { // left 是索引+1
		// 获取反转段的上一个节点
		p0 = p0.Next
	}

	// 同 反转链表 https://leetcode.cn/problems/reverse-linked-list/?envType=study-plan-v2&envId=top-100-liked
	var cur = p0.Next // 开始反转的起点
	var pre *ListNode = nil
	for i := 0; i < right-left+1; i++ {
		nxt := cur.Next
		cur.Next = pre

		pre = cur
		cur = nxt
	}

	// 循环结束性质: cur 指向下个循环段的开始,  pre指定末尾,返回 pre
	// 反转边界
	p0.Next.Next = cur
	p0.Next = pre // pre 这时为反转后的头节点

	return dummy.Next
}
