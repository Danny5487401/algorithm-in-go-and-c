package _2_link_list

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 从链表中移除节点 https://leetcode.cn/problems/remove-nodes-from-linked-list/description/
func TestRemoveNodes(t *testing.T) {
	convey.Convey("从链表中移除节点:保留右侧比左侧大的值", t, func() {
		testCase := []struct {
			input  *ListNode
			target *ListNode
		}{
			{
				getListNodeBySlice([]int{5, 2, 13, 3, 8}),
				getListNodeBySlice([]int{13, 8}),
			},
			{
				getListNodeBySlice([]int{1, 1, 1, 1}),
				getListNodeBySlice([]int{1, 1, 1, 1}),
			},
		}

		for _, tst := range testCase {
			rsp := removeNodes(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func removeNodes(head *ListNode) *ListNode {
	for head.Next == nil { // 1 <= Node.val <= 105
		return head
	}
	// 最大值，那么用递归解决，递归本质就是在倒着遍历链表
	node := removeNodes(head.Next) // 代表下个节点
	if node.Val > head.Val {       // 返回的链表头一定是最大的
		return node // 删除 head
	}
	head.Next = node // 不删除 head

	return head

}

// 方式二：通过 206. 反转链表，我们可以从反转后的链表头开始，像 83. 删除排序链表中的重复元素 那样，删除比当前节点值小的元素。最后再次反转链表，即为答案
