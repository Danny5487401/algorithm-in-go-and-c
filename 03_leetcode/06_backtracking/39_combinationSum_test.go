package _6_backtracking

import (
	"github.com/smartystreets/goconvey/convey"
	"slices"
	"testing"
)

// 组合总和 https://leetcode.cn/problems/combination-sum/description/?envType=study-plan-v2&envId=top-100-liked
func TestCombinationSum(t *testing.T) {
	convey.Convey("无限制重复被选取", t, func() {
		testCase := []struct {
			input  []int
			target int
			ans    [][]int
		}{
			{
				[]int{2, 3, 6, 7},
				7,
				[][]int{
					{2, 2, 3},
					{7},
				},
			},
			{
				[]int{2, 3, 5},
				8,
				[][]int{
					{2, 2, 2, 2},
					{7},
				},
			},
		}

		for _, tst := range testCase {
			rsp := combinationSum(tst.input, tst.target)
			convey.So(rsp, convey.ShouldEqual, tst.ans)
		}
	})

}

// 方法一：选或不选
func combinationSum(candidates []int, target int) [][]int {
	var ans = make([][]int, 0)
	var dfs func(i, target int)
	length := len(candidates)
	var path []int
	dfs = func(i, target int) {
		if i == length || target < 0 {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), path...))
			return
		}

		// 不选
		dfs(i+1, target)

		// 选
		path = append(path, candidates[i])
		dfs(i, target-candidates[i])
		path = path[:len(path)-1]

	}
	dfs(0, target)
	return ans
}

// 剪枝优化: 把 candidates 从小到大排序，如果递归中发现 left<candidates[i]，由于后面的数字只会更大，所以无法把 left 减小到 0，可以直接返回
func combinationSum2(candidates []int, target int) (ans [][]int) {
	slices.Sort(candidates)
	path := []int{}
	var dfs func(int, int)
	dfs = func(i, left int) {
		if left == 0 {
			// 找到一个合法组合
			ans = append(ans, slices.Clone(path))
			return
		}

		if i == len(candidates) || left < candidates[i] {
			return
		}

		// 不选
		dfs(i+1, left)

		// 选
		path = append(path, candidates[i])
		dfs(i, left-candidates[i])
		path = path[:len(path)-1] // 恢复现场
	}
	dfs(0, target)
	return ans
}
