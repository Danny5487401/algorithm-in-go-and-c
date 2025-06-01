package _2_link_list

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 反转偶数长度组的节点 https://leetcode.cn/problems/reverse-nodes-in-even-length-groups/
func TestReverseEvenLengthGroups(t *testing.T) {
	convey.Convey(" 反转偶数长度组的节点:"+
		"节点 1 分配给第一组\n"+
		"节点 2 和 3 分配给第二组\n"+
		"节点 4、5 和 6 分配给第三组，以此类推."+
		"最后一组的长度可能小于或者等于 1 + 倒数第二组的长度", t, func() {
		testCase := []struct {
			input  *ListNode
			target *ListNode
		}{
			{
				/*
					- 第一组长度为 1 ，奇数，没有发生反转。
					- 第二组长度为 2 ，偶数，节点反转。
					- 第三组长度为 3 ，奇数，没有发生反转。
					- 最后一组长度为 4 ，偶数，节点反转。
				*/
				GetListNodeBySlice([]int{5, 2, 6, 3, 9, 1, 7, 3, 8, 4}),
				GetListNodeBySlice([]int{5, 6, 2, 3, 9, 1, 4, 8, 3, 7}),
			},

			{
				GetListNodeBySlice([]int{1, 1, 0, 6}),
				GetListNodeBySlice([]int{1, 0, 1, 6}),
			},

			{
				// - 第一组长度为 1 ，没有发生反转。
				// - 最后一组长度为 1 ，没有发生反转。
				GetListNodeBySlice([]int{2, 1}),
				GetListNodeBySlice([]int{2, 1}),
			},
		}

		for _, tst := range testCase {
			rsp := reverseEvenLengthGroups(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func reverseEvenLengthGroups(head *ListNode) *ListNode {
	var nodes []*ListNode
	for node, size := head, 1; node != nil; node = node.Next {
		nodes = append(nodes, node)
		if len(nodes) == size || node.Next == nil { // 统计到 size 个节点，或到达链表末尾
			if n := len(nodes); n%2 == 0 { // 有偶数个节点
				for i := 0; i < n/2; i++ {
					nodes[i].Val, nodes[n-1-i].Val = nodes[n-1-i].Val, nodes[i].Val // 直接交换元素值
				}
			}
			nodes = nil
			size++
		}
	}
	return head
}
