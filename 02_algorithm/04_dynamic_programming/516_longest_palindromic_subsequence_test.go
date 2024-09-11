package _4_dynamic_programming

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 最长回文子序列 https://leetcode.cn/problems/longest-palindromic-subsequence/description/
func TestLongestPalindromeSubseq(t *testing.T) {
	convey.Convey("最长回文子序列 ", t, func() {
		testCase := []struct {
			input  string
			target int
		}{
			{
				"bbbab",
				4,
			},
			{
				"cbbd",
				2,
			},
		}

		for _, tst := range testCase {
			rsp := longestPalindromeSubseq2(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func longestPalindromeSubseq(s string) int {

	/*
			dfs(i,j)=
		- dfs(i+1,j-1)+2   s[i]=s[j]
		- max(dfs(i+1,j),dfs(i,j-1))  s[i]!=s[j]
	*/

	length := len(s)
	cache := make([][]int, length) // 外层是 i
	for i := range cache {
		cache[i] = make([]int, length) // target+0所以要加一
		for j := range cache[i] {
			cache[i][j] = -1 // -1 表示没用访问过
		}
	}

	var dfs func(i int, j int) int
	dfs = func(i int, j int) (res int) {
		if i > j { // 没有字母
			return 0
		}
		if i == j { // 一个字母
			return 1
		}
		if cache[i][j] != -1 {
			return cache[i][j]
		}
		defer func() {
			cache[i][j] = res
		}()

		if s[i] == s[j] {
			return dfs(i+1, j-1) + 2
		}
		return max(dfs(i+1, j), dfs(i, j-1))
	}

	return dfs(0, length-1)
}

func longestPalindromeSubseq2(s string) int {
	/*
		f[i][j] =
		- 0 i>j
		- 1 i=j
		- f[i+1][j-1]+2  s[i]=s[j]
		- max(f[i+1][j],f[i+1][j-1]) s[i]!=s[j]

		- 因为需要从 f[i+1]->f[i],所以 i 逆序，
		- 因为从 s[i][j-1]->f[i][j], 所以 j 正序
	*/
	length := len(s)
	f := make([][]int, length)
	for i := range f {
		f[i] = make([]int, length)
	}
	for i := length - 1; i >= 0; i-- {
		f[i][i] = 1
		for j := i + 1; j < length; j++ {
			if s[i] == s[j] {
				f[i][j] = f[i+1][j-1] + 2
			} else {
				f[i][j] = max(f[i+1][j], f[i][j-1])
			}
		}
	}
	return f[0][length-1]
}
