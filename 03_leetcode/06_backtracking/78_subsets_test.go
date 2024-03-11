package _6_backtracking

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 子集 https://leetcode.cn/problems/subsets/?envType=study-plan-v2&envId=top-100-liked

func TestSubsets(t *testing.T) {
	convey.Convey("subsets ", t, func() {
		testCase := []struct {
			input  []int
			target [][]int
		}{
			{

				[]int{1, 2, 3}, [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}},
			},
		}

		for _, tst := range testCase {
			rsp := subsets(tst.input)
			convey.So(rsp, convey.ShouldResemble, tst.target)
		}
	})

}

// 方法一：输入的视角（选或不选）
func subsets(nums []int) [][]int {
	length := len(nums)
	if length == 0 {
		return [][]int{}
	}
	var ans = make([][]int, 0)
	var path = make([]int, 0, length)
	var dfs func(i int)
	dfs = func(i int) {
		if i == length {
			ans = append(ans, append([]int(nil), path...)) // 固定答案
			return
		}
		// 1 不选
		dfs(i + 1)
		// 2 选
		path = append(path, nums[i])
		dfs(i + 1)

		// 返回恢复现场
		path = path[:len(path)-1]
	}
	dfs(0)
	return ans
}
