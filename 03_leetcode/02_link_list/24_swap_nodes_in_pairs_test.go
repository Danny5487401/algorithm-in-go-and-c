package _2_link_list

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 两两交换链表中的节点 https://leetcode.cn/problems/swap-nodes-in-pairs/description/
func TestSwapPairs(t *testing.T) {
	convey.Convey("两两交换链表中的节点:两两交换其中相邻的节点 ", t, func() {
		testCase := []struct {
			input    *ListNode
			expected *ListNode
		}{
			{

				GetListNodeBySlice([]int{1, 2, 3, 4}), GetListNodeBySlice([]int{2, 1, 4, 3}),
			},
			{
				GetListNodeBySlice([]int{}), GetListNodeBySlice([]int{}),
			},
			{
				GetListNodeBySlice([]int{1}), GetListNodeBySlice([]int{1}),
			},
		}

		for _, tst := range testCase {
			rsp := swapPairs(tst.input)
			convey.So(rsp, convey.ShouldResemble, tst.expected)
		}
	})

}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head} // 用哨兵节点简化代码逻辑

	node0 := dummy
	node1 := head
	for node1 != nil && node1.Next != nil { // 至少有两个节点

		// node0 -> node1 -> node2 -> node3

		// 记录下次开始节点 node2 node3
		node2 := node1.Next
		node3 := node2.Next

		// node0 -> node2 -> node1 -> node3
		node0.Next = node2 // 0 -> 2
		node2.Next = node1 // 2 -> 1
		node1.Next = node3 // 1 -> 3

		node0 = node1 // 下一轮交换，0 是 1
		node1 = node3 // 下一轮交换，1 是 3
	}
	return dummy.Next // 返回新链表的头节点
}

// 方式二： 使用 K 个一组翻转链表 方式  https://leetcode.cn/problems/reverse-nodes-in-k-group/?envType=study-plan-v2&envId=top-100-liked
func swapPairsMethod2(head *ListNode) *ListNode {
	var n = 0 // 链表长度
	curTmp := head
	for curTmp != nil {
		n++ // 获取链表长度
		curTmp = curTmp.Next
	}

	var dummy = &ListNode{
		Next: head,
	}
	p0 := dummy // 反转段的上一个节点
	for n >= 2 {
		n -= 2
		// 同反转中间部分链表 https://leetcode.cn/problems/reverse-linked-list-ii/description/

		var pre *ListNode = nil
		var cur = p0.Next // 开始反转的起点
		for i := 0; i < 2; i++ {
			nxt := cur.Next
			cur.Next = pre
			pre = cur
			cur = nxt
		}

		// 记录下次循环的开始
		nxt := p0.Next
		p0.Next.Next = cur
		p0.Next = pre
		p0 = nxt
	}

	return dummy.Next
}

// GetListNodeBySlice 通过 slice 构建 list
func GetListNodeBySlice(s []int) *ListNode {
	dummy := &ListNode{}
	head := dummy

	for _, v := range s {
		head.Next = &ListNode{Val: v}
		head = head.Next
	}
	return dummy.Next
}
