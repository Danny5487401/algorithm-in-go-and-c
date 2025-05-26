package _2_link_list

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 合并零之间的节点 https://leetcode.cn/problems/merge-nodes-in-between-zeros/description/
func TestMrgeNodes(t *testing.T) {
	convey.Convey("合并零之间的节点: 对于每两个相邻的 0 ，请你将它们之间的所有节点合并成一个节点，其值是所有已合并节点的值之和。", t, func() {
		testCase := []struct {
			input  *ListNode
			target *ListNode
		}{
			{
				// 可以把 3,1 的和填到第一个节点 0 上, 把 4,5,2 的和填到第二个节点 3 上。
				getListNodeBySlice([]int{0, 3, 1, 0, 4, 5, 2, 0}),
				getListNodeBySlice([]int{4, 11}),
			},
		}

		for _, tst := range testCase {
			rsp := mergeNodes(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})
}

// 方法一: 两条链
func mergeNodes(head *ListNode) *ListNode {
	// 不 存在连续两个 Node.val == 0 的节点
	// 链表的 开端 和 末尾 节点都满足 Node.val == 0
	var dummyNode = &ListNode{}
	var cur = dummyNode
	temp := 0
	for head != nil {
		nxt := head.Next
		if head.Val == 0 {
			newNode := &ListNode{
				Val: temp,
			}
			cur.Next = newNode
			cur = cur.Next
			head = nxt
			temp = 0
			continue
		}
		temp += head.Val
		head = nxt
	}
	return dummyNode.Next.Next
}

// 方法二: 一条链

func mergeNodes2(head *ListNode) *ListNode {
	tail := head //  head 当作答案链表的头节点。在一开始，head 也是答案链表的末尾节点 tail
	for cur := head.Next; cur.Next != nil; cur = cur.Next {
		if cur.Val != 0 {
			tail.Val += cur.Val
		} else {
			tail = tail.Next
			tail.Val = 0 // 重新归零
		}
	}
	tail.Next = nil // 由于我们是在原链表上修改的，原链表的最后一个节点 0 必然不在答案链表中，所以答案链表一定比原链表短。把 tail.next 置为空，以确保答案链表不包含原链表中多余的节点

	return head
}
