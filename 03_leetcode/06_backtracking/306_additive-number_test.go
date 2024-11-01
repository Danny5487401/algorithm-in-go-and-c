package _6_backtracking

import (
	"github.com/smartystreets/goconvey/convey"
	"strconv"
	"testing"
)

// 累加数: https://leetcode.cn/problems/additive-number/description/
func TestIsAdditiveNumber(t *testing.T) {
	convey.Convey("累加数:除了最开始的两个数以外，序列中的每个后续数字必须是它之前两个数字之和", t, func() {
		testCase := []struct {
			input  string
			target bool
		}{
			{
				"112358",
				true,
			},
			{
				"199100199",
				true,
			},
		}

		for _, tst := range testCase {
			rsp := isAdditiveNumber(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

// 输入视角(选/不选)
func isAdditiveNumber(num string) bool {
	n := len(num)
	var path []int
	var ans [][]int // 列出所有组合
	var dfs func(int, int)
	dfs = func(i, start int) {
		if i == n {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		// 不选
		if i < n-1 {
			dfs(i+1, start)
		}
		if i-start > 0 && num[start] == '0' {
			return
		}
		digit, _ := strconv.Atoi(num[start : i+1])
		// 选
		if len(path) < 2 || digit == path[len(path)-1]+path[len(path)-2] {
			path = append(path, digit)
			dfs(i+1, i+1)
			path = path[:len(path)-1]
		}

	}
	dfs(0, 0)

	flag := false
	for _, array := range ans {
		if len(array) >= 3 { // 代表有组合
			flag = true
			break
		}
	}
	return flag
}
