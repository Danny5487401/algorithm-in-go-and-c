package _8_array

import (
	"github.com/smartystreets/goconvey/convey"
	"slices"
	"testing"
)

// 合并区间 https://leetcode.cn/problems/merge-intervals/description/?envType=study-plan-v2&envId=top-100-liked
func TestMerge(t *testing.T) {
	convey.Convey("合并区间", t, func() {
		testCase := []struct {
			input  [][]int
			target [][]int
		}{
			{

				[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}},
			},

			{

				[][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}},
			},
		}

		for _, tst := range testCase {
			rsp := merge(tst.input)
			convey.So(rsp, convey.ShouldResemble, tst.target)
		}
	})

}

func merge(intervals [][]int) (ans [][]int) {
	// 按照左端点从小到大排序
	slices.SortFunc(intervals, func(i, j []int) int { return i[0] - j[0] })

	for _, p := range intervals {
		m := len(ans)
		if m > 0 && p[0] <= ans[m-1][1] { // 遍历左端点小于答案的右端点
			// 合并更新右端点
			ans[m-1][1] = p[1]
		} else {
			ans = append(ans, p)
		}
	}
	return
}
