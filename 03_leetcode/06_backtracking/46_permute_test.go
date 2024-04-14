package _6_backtracking

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 全排列 https://leetcode.cn/problems/permutations/description/

func TestPermute(t *testing.T) {
	convey.Convey("permute 全排列 ", t, func() {
		testCase := []struct {
			input  []int
			target [][]int
		}{
			{

				[]int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
			},
			{

				[]int{0, 1}, [][]int{{0, 1}, {1, 0}},
			},
			{

				[]int{1}, [][]int{{1}},
			},
		}

		for _, tst := range testCase {
			rsp := permute(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func permute(nums []int) [][]int {
	var ans = make([][]int, 0)
	length := len(nums)
	// 是否已选
	onPath := make([]bool, length)
	var dfs func(i int)
	var path = make([]int, length)
	dfs = func(i int) {
		if i == length {
			ans = append(ans, append([]int(nil), path...)) // 固定答案
			return
		}
		for j, on := range onPath {
			if !on {
				// 没有选的信息
				path[i] = nums[j]
				onPath[j] = true
				dfs(i + 1)
				onPath[j] = false
			}
		}

	}
	dfs(0)
	return ans
}
