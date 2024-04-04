package _4_dynamic_programming

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 最长公共子序列 https://leetcode.cn/problems/longest-common-subsequence/

func TestLongestCommonSubsequence(t *testing.T) {
	convey.Convey("最长公共子序列 ", t, func() {
		testCase := []struct {
			input1 string
			input2 string
			target int
		}{
			{

				"abcde", "ace", 3,
			},
			{

				"abc", "abc", 3,
			},
			{

				"abc", "def", 0,
			},
		}

		for _, tst := range testCase {
			rsp := longestCommonSubsequence2(tst.input1, tst.input2)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func longestCommonSubsequence(text1 string, text2 string) int {
	// 方案一： 递归搜索 + 保存计算结果 = 记忆化搜索
	n := len(text1)
	m := len(text2)
	var cache = make([][]int, n)
	for i := range cache {
		cache[i] = make([]int, m)
		for j := range cache[i] {
			cache[i][j] = -1 // -1 表示没有访问过
		}

	}

	var dfs func(i, j int) (res int)
	dfs = func(i, j int) (res int) {
		if i < 0 || j < 0 {
			return
		}
		c := cache[i][j]
		if c != -1 {
			return cache[i][j]
		}
		defer func() {
			cache[i][j] = res
		}()
		if text1[i] == text2[j] {
			return dfs(i-1, j-1) + 1
		}
		return max(dfs(i-1, j), dfs(i, j-1))
	}
	return dfs(n-1, m-1)
}

// 1:1 翻译成递推

func longestCommonSubsequence2(text1 string, text2 string) int {
	n, m := len(text1), len(text2)

	// f[i][j] = max(f[i-1, j], dfs(i, j-1))-->f[i+1][j+1] =max(f[i][j+1], f[i+1][j])
	// f[i][j] = dfs(i-1, j-1) + 1 -->f[i+1][j+1]=f[i][j]+1
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
	}
	for i, x := range text1 {
		for j, y := range text2 {
			if x == y {
				f[i+1][j+1] = f[i][j] + 1
			} else {
				f[i+1][j+1] = max(f[i][j+1], f[i+1][j])
			}
		}
	}

	return f[n][m]
}
