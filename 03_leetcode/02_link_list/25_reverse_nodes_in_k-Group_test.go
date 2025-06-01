package _2_link_list

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// K 个一组翻转链表 https://leetcode.cn/problems/reverse-nodes-in-k-group/?envType=study-plan-v2&envId=top-100-liked
func TestReverseKGroup(t *testing.T) {
	convey.Convey(" K 个一组翻转链表:如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序", t, func() {
		testCase := []struct {
			input  *ListNode
			k      int
			target *ListNode
		}{
			{
				GetListNodeBySlice([]int{1, 2, 3, 4, 5}),
				2,
				GetListNodeBySlice([]int{2, 1, 4, 3, 5}),
			},

			{
				GetListNodeBySlice([]int{1, 2, 3, 4, 5}),
				3,
				GetListNodeBySlice([]int{3, 2, 1, 4, 5}),
			},
		}

		for _, tst := range testCase {
			rsp := reverseKGroup(tst.input, tst.k)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func reverseKGroup(head *ListNode, k int) *ListNode {
	var n = 0 // 链表长度
	curTmp := head
	for curTmp != nil {
		n++ // 获取链表长度
		curTmp = curTmp.Next
	}

	// 创建哨兵节点
	var dummy = &ListNode{
		Next: head,
	}
	p0 := dummy // 反转段的上一个节点
	for n >= k {
		n -= k

		// 记录下次循环的开始
		nxtRound := p0.Next

		// 同反转中间部分链表 https://leetcode.cn/problems/reverse-linked-list-ii/description/

		var pre *ListNode = nil
		var cur = p0.Next // 开始反转的起点
		for i := 0; i < k; i++ {
			nxt := cur.Next
			cur.Next = pre
			pre = cur
			cur = nxt
		}

		// 循环结束性质: cur 指向下个循环段的开始,  pre指定末尾,返回 pre
		// 反转边界
		p0.Next.Next = cur
		p0.Next = pre

		// 更改反转段的上一个节点
		p0 = nxtRound
	}

	return dummy.Next
}
