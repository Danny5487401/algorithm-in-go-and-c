package _6_backtracking

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

// 组合 https://leetcode.cn/problems/combinations/description/ 是 子集的特殊情况 https://leetcode.cn/problems/subsets/?envType=study-plan-v2&envId=top-100-liked

func TestCombination(t *testing.T) {
	convey.Convey("组合 ", t, func() {
		testCase := []struct {
			input struct {
				n int
				k int
			}
			target [][]int
		}{
			{

				struct {
					n int
					k int
				}{n: 4, k: 2}, [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
			},
			{

				struct {
					n int
					k int
				}{n: 1, k: 1}, [][]int{{1}},
			},
		}

		for _, tst := range testCase {
			rsp := combine(tst.input.n, tst.input.k)
			convey.So(rsp, convey.ShouldResemble, tst.target)
		}
	})

}

func combine(n int, k int) [][]int {
	var ans = make([][]int, 0)

	var dfs func(i int)
	var path = make([]int, 0)
	dfs = func(i int) {

		d := k - len(path) // 还要选 d 个数
		if d == 0 {
			ans = append(ans, append([]int(nil), path...))
			return
		}

		// 从后往前容易点
		for j := i; j >= d; j-- {
			path = append(path, j)
			dfs(j - 1)
			// 返回恢复现场
			path = path[:len(path)-1]
		}
	}
	dfs(n)
	return ans
}
