package _2_link_list

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 从链表中移除节点 https://leetcode.cn/problems/remove-nodes-from-linked-list/description/
func TestRemoveNodes(t *testing.T) {
	convey.Convey("从链表中移除节点:移除每个右侧有一个更大数值的节点", t, func() {
		testCase := []struct {
			input  *ListNode
			target *ListNode
		}{
			{
				/*
					需要移除的节点是 5 ，2 和 3 。
					- 节点 13 在节点 5 右侧。
					- 节点 13 在节点 2 右侧。
					- 节点 8 在节点 3 右侧。
				*/
				getListNodeBySlice([]int{5, 2, 13, 3, 8}),
				getListNodeBySlice([]int{13, 8}),
			},
			{
				// 每个节点的值都是 1 ，所以没有需要移除的节点。
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

// 思路: https://leetcode.cn/problems/remove-nodes-from-linked-list/solutions/1993491/di-gui-jian-ji-xie-fa-by-endlesscheng-jfwi/
func removeNodes(head *ListNode) *ListNode {
	// 1 <= Node.val <= 10^5
	for head.Next == nil {
		return head
	}
	// 最大值，那么用递归解决，递归本质就是在倒着遍历链表
	// 先对它的右侧节点进行移除操作，得到一个新的子链表，如果子链表的表头节点值大于该节点的值，那么移除该节点
	node := removeNodes(head.Next) // 代表下个节点
	if node.Val > head.Val {       // 返回的链表头一定是最大的
		return node // 删除 head
	}
	// 将该节点作为子链表的表头节点，最后返回该子链表。
	head.Next = node // 不删除 head

	return head

}

// 方式二：通过 206. 反转链表，我们可以从反转后的链表头开始，像 8->3. 删除排序链表中的重复元素 那样，删除比当前节点值小的元素。最后再次反转链表，即为答案
func removeNodes2(head *ListNode) *ListNode {
	head = reverseList(head)
	cur := head
	for cur.Next != nil {
		if cur.Val > cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return reverseList(head)
}
