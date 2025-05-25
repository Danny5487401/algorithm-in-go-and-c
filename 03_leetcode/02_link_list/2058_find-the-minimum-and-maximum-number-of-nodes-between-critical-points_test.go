package _2_link_list

import (
	"github.com/smartystreets/goconvey/convey"
	"math"
	"testing"
)

// 找出临界点之间的最小和最大距离 https://leetcode.cn/problems/find-the-minimum-and-maximum-number-of-nodes-between-critical-points/description/
func TestNodesBetweenCriticalPoints(t *testing.T) {
	convey.Convey("找出临界点之间的最小和最大距离: "+
		"如果当前节点的值 严格大于 前一个节点和后一个节点，那么这个节点就是一个  局部极大值点 。"+
		"如果当前节点的值 严格小于 前一个节点和后一个节点，那么这个节点就是一个  局部极小值点 。", t, func() {
		testCase := []struct {
			input  *ListNode
			target []int
		}{
			{
				getListNodeBySlice([]int{3, 1}),
				[]int{-1, -1},
			},
			{
				/*
					- [5,3,1,2,5,1,2]：第三个节点是一个局部极小值点，因为 1 比 3 和 2 小。
					- [5,3,1,2,5,1,2]：第五个节点是一个局部极大值点，因为 5 比 2 和 1 大。
					- [5,3,1,2,5,1,2]：第六个节点是一个局部极小值点，因为 1 比 5 和 2 小。
					第五个节点和第六个节点之间距离最小。minDistance = 6 - 5 = 1 。
					第三个节点和第六个节点之间距离最大。maxDistance = 6 - 3 = 3
				*/
				getListNodeBySlice([]int{5, 3, 1, 2, 5, 1, 2}),
				[]int{1, 3},
			},
			{
				/*
					- [1,3,2,2,3,2,2,2,7]：第二个节点是一个局部极大值点，因为 3 比 1 和 2 大。
					- [1,3,2,2,3,2,2,2,7]：第五个节点是一个局部极大值点，因为 3 比 2 和 2 大。
					最小和最大距离都存在于第二个节点和第五个节点之间。
					因此，minDistance 和 maxDistance 是 5 - 2 = 3 。
				*/
				getListNodeBySlice([]int{1, 3, 2, 2, 3, 2, 2, 2, 7}),
				[]int{3, 3},
			},
			{
				getListNodeBySlice([]int{2, 3, 3, 2}),
				[]int{-1, -1},
			},
		}

		for _, tst := range testCase {
			rsp := nodesBetweenCriticalPoints(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func nodesBetweenCriticalPoints(head *ListNode) []int {
	// 链表中节点的数量在范围 [2, 105] 内
	a, b, c := head, head.Next, head.Next.Next
	first, last, minDis := 0, 0, math.MaxInt32
	for i, prev := 1, 0; c != nil; i++ { // 遍历链表，寻找临界点
		if a.Val < b.Val && b.Val > c.Val || a.Val > b.Val && b.Val < c.Val {
			if first == 0 {
				first = i // 首个临界点位置
			}
			last = i // 最末临界点位置
			if prev > 0 && i-prev < minDis {
				minDis = i - prev // 更新相邻临界点位置之差的最小值
			}
			prev = i
		}
		a, b, c = b, c, c.Next
	}
	if first == last { // 临界点少于两个
		return []int{-1, -1}
	}
	return []int{minDis, last - first}
}
