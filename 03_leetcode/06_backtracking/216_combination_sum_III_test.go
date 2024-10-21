package _6_backtracking

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 组合总和III :https://leetcode.cn/problems/combination-sum-iii/solutions/2071013/hui-su-bu-hui-xie-tao-lu-zai-ci-pythonja-feme/

func TestCombination3(t *testing.T) {
	convey.Convey("组合总和 III:相加之和为 n 的 k 个数的组合,只使用数字1到9,每个数字 最多使用一次  ", t, func() {
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
			rsp := combinationSumIII_1(tst.input.k, tst.input.n)
			compareRsp := intSliceSliceEqual(rsp, tst.target)
			convey.So(compareRsp, convey.ShouldBeTrue)
		}
	})

}

// 方式二：多几个剪枝
// 元素和超过 n 优化
// 元素和选择最大前几个小于 n 优化

func combinationSumIII_2(k int, n int) [][]int {
	var ans = make([][]int, 0)

	var dfs func(i, t int)

	var path = make([]int, 0)

	dfs = func(i, t int) { // t 代表目标值
		m := len(path)
		d := k - m // 还要选 d 个数

		// 这里进行剪枝优化
		if t < 0 || t > (i+i-d+1)*d/2 { // 剩余选择最大的
			return
		}

		if d == 0 { // 上面可以判断 t<0 || t>0
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

// 方法一：枚举下一个数选哪个
func combinationSumIII_1(k int, n int) [][]int {
	var ans = make([][]int, 0)

	var dfs func(i int)

	var path = make([]int, 0)

	dfs = func(i int) {
		m := len(path)
		d := k - m // 还要选 d 个数

		if d == 0 {
			// 相比 77 组合 https://leetcode.cn/problems/combinations/description/ 简单多个判断
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
