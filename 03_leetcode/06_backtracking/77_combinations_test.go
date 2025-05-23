package _6_backtracking

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

// 组合 https://leetcode.cn/problems/combinations/description/ 是 子集的特殊情况 https://leetcode.cn/problems/subsets/?envType=study-plan-v2&envId=top-100-liked
func TestCombination(t *testing.T) {
	convey.Convey("组合: [1,n] 选 k 个数", t, func() {
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
			compareRsp := intSliceSliceEqual(rsp, tst.target)
			convey.So(compareRsp, convey.ShouldBeTrue)
		}
	})

}

// 方法一：枚举下一个数选哪个
func combine(n int, k int) [][]int {
	var ans = make([][]int, 0)
	var dfs func(i int)
	var path = make([]int, 0)
	dfs = func(i int) { // i 代表范围 [0,i]
		m := len(path) // m 代表已经选择的数量
		d := k - m     // 还要选 d 个数
		if d == 0 {
			ans = append(ans, append([]int(nil), path...)) // 记录答案
			return
		}

		for j := i; j >= d; j-- { // 这里的意思就是 d>=i
			path = append(path, j)
			dfs(j - 1)
			// 返回恢复现场
			path = path[:len(path)-1]
		}
	}
	dfs(n) // 倒序容易点：从后往前
	return ans
}

// 方法二：选或不选
func combine2(n, k int) (ans [][]int) {
	path := []int{}
	var dfs func(int)
	dfs = func(i int) {
		d := k - len(path) // 还要选 d 个数
		if d == 0 {        // 选好了
			ans = append(ans, append([]int(nil), path...))
			return
		}

		// 如果 i > d，可以不选 i
		if i > d {
			dfs(i - 1) // 不选 i
		}

		// 选 i
		path = append(path, i)
		dfs(i - 1)
		path = path[:len(path)-1] // 恢复现场
	}
	dfs(n)
	return
}
