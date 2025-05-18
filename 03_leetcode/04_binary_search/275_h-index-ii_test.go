package _4_binary_search

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// H 指数 II: https://leetcode.cn/problems/h-index-ii/description/
func TestHIndexII(t *testing.T) {
	convey.Convey("H 指数 II:citations 已经按照 非降序排列 , h 代表“高引用次数”（high citations），一名科研人员的 h 指数是指他（她）的 （n 篇论文中）至少 有 h 篇论文分别被引用了至少 h 次", t, func() {
		testCase := []struct {
			input    []int
			expected int
		}{
			{
				// 给定数组表示研究者总共有 5 篇论文，每篇论文相应的被引用了 0, 1, 3, 5, 6 次。
				// 由于研究者有3篇论文每篇 至少 被引用了 3 次，其余两篇论文每篇被引用 不多于 3 次，所以她的 h 指数是 3
				[]int{0, 1, 3, 5, 6}, 3,
			},
			{
				[]int{1, 2, 100}, 2,
			},
		}

		for _, tst := range testCase {
			rsp := hIndex(tst.input)
			convey.So(rsp, convey.ShouldResemble, tst.expected)
		}
	})

}

func hIndex(citations []int) int {

	// 解题思路: https://leetcode.cn/problems/h-index-ii/solutions/2504326/tu-jie-yi-tu-zhang-wo-er-fen-da-an-si-ch-d15k/
	// 在区间 [left, right] 内询问引用的次数
	left, right := 1, len(citations)
	for left <= right { // 区间不为空
		// 循环不变量：
		// left-1 的回答一定为「是」
		// right+1 的回答一定为「否」
		mid := (left + right) / 2
		// 引用次数最多的 mid 篇论文，引用次数均 >= mid
		if citations[len(citations)-mid] >= mid {
			left = mid + 1 // 询问范围缩小到 [mid+1, right]
		} else {
			right = mid - 1 // 询问范围缩小到 [left, mid-1]
		}
	}
	// 循环结束后 right 等于 left-1，回答一定为「是」
	// 根据循环不变量，right 现在是最大的回答为「是」的数
	return right

}
