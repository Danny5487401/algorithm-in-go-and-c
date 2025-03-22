package _2_link_list

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 奇偶链表 https://leetcode.cn/problems/odd-even-linked-list/description/
func TestOddEvenList(t *testing.T) {
	convey.Convey("将所有索引为奇数的节点和索引为偶数的节点分别组合在一起, O(1) 的额外空间复杂度", t, func() {
		testCase := []struct {
			input    *ListNode
			expected *ListNode
		}{
			{

				GetListNodeBySlice([]int{1, 2, 3, 4, 5}), GetListNodeBySlice([]int{1, 3, 5, 2, 4}),
			},
			{
				GetListNodeBySlice([]int{2, 1, 3, 5, 6, 4, 7}), GetListNodeBySlice([]int{2, 3, 6, 7, 1, 5, 4}),
			},
		}

		for _, tst := range testCase {
			rsp := oddEvenList(tst.input)
			convey.So(rsp, convey.ShouldResemble, tst.expected)
		}
	})
}

// 分离节点后合并
func oddEvenList(head *ListNode) *ListNode {

	/*

	   原始链表的头节点 head 也是奇数链表的头节点以及结果链表的头节点.令 evenHead = head.next，则 evenHead 是偶数链表的头节点
	*/
	if head == nil {
		return head
	}
	// 头节点是奇数节点，头节点的后一个节点是偶数节点，相邻节点的奇偶性不同。
	evenHead := head.Next // 分离偶数节点

	odd := head      // 奇数数索引节点
	even := evenHead // 偶数索引节点
	for even != nil && even.Next != nil {
		odd.Next = even.Next // 更新奇数节点时，奇数节点的后一个节点需要指向偶数节点的后一个节点
		odd = odd.Next

		even.Next = odd.Next // 更新偶数节点时，偶数节点的后一个节点需要指向奇数节点的后一个节点
		even = even.Next
	}
	odd.Next = evenHead // 合并
	return head

}
