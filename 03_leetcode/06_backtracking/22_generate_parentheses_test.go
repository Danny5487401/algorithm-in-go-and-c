package _6_backtracking

import (
	"github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

// 括号生成 https://leetcode.cn/problems/generate-parentheses/

func TestGenerateParenthesis(t *testing.T) {
	convey.Convey("括号生成 ", t, func() {
		testCase := []struct {
			input struct {
				n int
			}
			target []string
		}{
			{

				struct {
					n int
				}{n: 3}, []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
			},
			{

				struct {
					n int
				}{n: 1}, []string{"()"},
			},
		}

		for _, tst := range testCase {
			rsp := generateParenthesis(tst.input.n)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func generateParenthesis(n int) []string {
	total := n * 2

	ans := make([]string, 0)

	path := make([]string, total)
	var dfs func(i int, left int)

	dfs = func(i int, left int) {
		if i == total {
			ans = append(ans, strings.Join(path, ""))
		}

		if left < n {
			// 说明还可以选择左括号
			path[i] = "("
			dfs(i+1, left+1)
		}

		if i-left < left {
			// 右括号小于左括号
			path[i] = ")"
			dfs(i+1, left)
		}
	}

	dfs(0, 0)
	return ans

}
