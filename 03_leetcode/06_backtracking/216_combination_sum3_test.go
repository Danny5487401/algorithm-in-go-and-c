package _6_backtracking

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 组合总和III :https://leetcode.cn/problems/combination-sum-iii/solutions/2071013/hui-su-bu-hui-xie-tao-lu-zai-ci-pythonja-feme/

func TestCombination3(t *testing.T) {
	convey.Convey("组合总和III ", t, func() {
		testCase := []struct {
			input struct {
				k int
				n int
			}
			target [][]int
		}{
			{

				struct {
					k int
					n int
				}{n: 7, k: 3}, [][]int{{1, 2, 4}},
			},
			{

				struct {
					k int
					n int
				}{n: 9, k: 3}, [][]int{
					{1, 2, 6},
					{1, 3, 5},
					{2, 3, 4},
				},
			},
			{

				struct {
					k int
					n int
				}{n: 1, k: 4}, [][]int{},
			},
		}

		for _, tst := range testCase {
			rsp := combinationSum3(tst.input.k, tst.input.n)
			compareRsp := intSliceSliceEqual(rsp, tst.target)
			convey.So(compareRsp, convey.ShouldBeTrue)
		}
	})

}

// 方式二：多几个剪枝
// 元素和超过 n 优化
// 元素和选择最大前几个小于 n 优化

func combinationSum3(k int, n int) [][]int {
	var ans = make([][]int, 0)

	var dfs func(i, t int)

	var path = make([]int, 0)

	dfs = func(i, t int) {
		m := len(path)
		d := k - m // 还要选 d 个数

		if t < 0 || t > (i*2-d+1)*d/2 {
			return
		}

		if d == 0 { // 上面可以判断 t<0 ||t>0

			ans = append(ans, append([]int(nil), path...)) // 记录答案

			return
		}

		for j := i; j >= d; j-- { // 这里的意思就是 d>=i
			path = append(path, j)
			dfs(j-1, t-j)
			// 返回恢复现场
			path = path[:len(path)-1]
		}

	}

	dfs(9, n)
	return ans
}

// 方式一：相比77 简单多个判断
func combinationSum3Extra1(k int, n int) [][]int {
	var ans = make([][]int, 0)

	var dfs func(i int)

	var path = make([]int, 0)

	dfs = func(i int) {
		m := len(path)
		d := k - m // 还要选 d 个数

		if d == 0 {
			var sum int
			for _, v := range path {
				sum += v
			}
			if sum == n {
				ans = append(ans, append([]int(nil), path...)) // 记录答案
			}

			return
		}

		for j := i; j >= d; j-- { // 这里的意思就是 d>=i
			path = append(path, j)
			dfs(j - 1)
			// 返回恢复现场
			path = path[:len(path)-1]
		}

	}

	dfs(9)
	return ans
}
