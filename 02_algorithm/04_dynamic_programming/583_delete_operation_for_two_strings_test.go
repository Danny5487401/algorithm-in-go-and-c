package _4_dynamic_programming

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 两个字符串的删除操作 https://leetcode.cn/problems/delete-operation-for-two-strings/description/
func TestMinDistance2(t *testing.T) {
	convey.Convey("两个字符串的删除操作", t, func() {
		testCase := []struct {
			input1 string
			input2 string
			target int
		}{
			{
				// 第一步将 "sea" 变为 "ea" ，第二步将 "eat "变为 "ea"
				"sea", "eat", 2,
			},
			{

				"leetcode", "etco", 4,
			},
		}

		for _, tst := range testCase {
			rsp := minDistance2(tst.input1, tst.input2)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func minDistance2(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)

	var cache = make([][]int, m)
	for i := range cache {
		cache[i] = make([]int, n)
		for j := range cache[i] {
			cache[i][j] = -1 // -1 表示没用访问过
		}
	}
	var dfs func(i int, j int) int
	dfs = func(i int, j int) (res int) {
		if i < 0 {
			return j + 1
		}
		if j < 0 {
			return i + 1
		}
		if cache[i][j] != -1 {
			return cache[i][j]
		}
		defer func() {
			cache[i][j] = res
		}()
		if word1[i] == word2[j] {
			return dfs(i-1, j-1) // 相等，不需要操作
		}
		return min(dfs(i, j-1), dfs(i-1, j)) + 1
	}

	return dfs(m-1, n-1)
}
