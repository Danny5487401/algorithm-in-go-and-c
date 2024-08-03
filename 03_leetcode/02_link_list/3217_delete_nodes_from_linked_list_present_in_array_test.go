package _2_link_list

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 从链表中移除在数组中存在的节点 https://leetcode.cn/problems/delete-nodes-from-linked-list-present-in-array/description/
func TestModifiedList(t *testing.T) {
	convey.Convey("从链表中移除在数组中存在的节点", t, func() {
		testCase := []struct {
			input  *ListNode
			nums   []int
			target *ListNode
		}{
			{
				getListNodeBySlice([]int{1, 2, 3, 4, 5}),
				[]int{1, 2, 3},
				getListNodeBySlice([]int{4, 5}),
			},
			{
				getListNodeBySlice([]int{1, 2, 1, 2, 1, 2}),
				[]int{1},
				getListNodeBySlice([]int{2, 2, 2}),
			},
			{
				getListNodeBySlice([]int{1, 2, 3, 4}),
				[]int{5},
				getListNodeBySlice([]int{1, 2, 3, 4}),
			},
		}

		for _, tst := range testCase {
			rsp := modifiedList(tst.nums, tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func modifiedList(nums []int, head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	// 由于直接判断节点值是否在 nums 中，需要遍历 nums，时间复杂度为 O(n)。
	// 通过把 nums 中的元素加到一个哈希集合中，然后判断节点值是否在哈希集合中，这样可以做到每次判断时间复杂度为 O(1)
	deleteMap := make(map[int]struct{})
	for _, n := range nums {
		deleteMap[n] = struct{}{}
	}

	var dummyNode = &ListNode{Next: head}

	cur := dummyNode

	for cur.Next != nil {
		if _, ok := deleteMap[cur.Next.Val]; ok {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummyNode.Next

}
