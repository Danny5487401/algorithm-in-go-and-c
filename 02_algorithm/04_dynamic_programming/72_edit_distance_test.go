package _4_dynamic_programming

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 编辑距离 https://leetcode.cn/problems/edit-distance/description/
func TestMinDistance(t *testing.T) {
	convey.Convey("编辑距离:增删改字符 ", t, func() {
		testCase := []struct {
			input1 string
			input2 string
			target int
		}{
			{

				"horse", "ros", 3,
			},
			{

				"intention", "execution", 5,
			},
		}

		for _, tst := range testCase {
			rsp := minDistance(tst.input1, tst.input2)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func minDistance(word1 string, word2 string) int {
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
			return dfs(i-1, j-1)
		}
		// 分别对应 新增，删除，替换
		return min(dfs(i, j-1), dfs(i-1, j), dfs(i-1, j-1)) + 1
	}

	return dfs(m-1, n-1)
}
