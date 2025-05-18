package _4_dynamic_programming

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 编辑距离 https://leetcode.cn/problems/edit-distance/description/
func TestMinDistance(t *testing.T) {
	convey.Convey("编辑距离:增删改字符,请返回将 word1 转换成 word2 所使用的最少操作数 , ", t, func() {
		testCase := []struct {
			input1 string
			input2 string
			target int
		}{
			{
				/*
					horse -> rorse (replace 'h' with 'r')
					rorse -> rose (remove 'r')
					rose -> ros (remove 'e')
				*/
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
	/*
		dfs(i,j)=
		  - s[i]=s[j]: dfs[i-1,j-1]                                相同
		  - s[i]!=s[j]: min(dfs(i,j-1),dfs(i-1,j),dfs(i-1,j-1))+1  插入,删除,替换
	*/
	m := len(word1)
	n := len(word2)

	var cache = make([][]int, m)
	for i := range cache {
		cache[i] = make([]int, n)
		for j := range cache[i] {
			cache[i][j] = -1 // -1 表示没用访问过
		}
	}
	var dfs func(i int, j int) int // 假设两个字符的下标为 i 和 j，先操作 i 再操作 j，还是先操作 j 再操作 i，结果都是一样的
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
			return dfs(i-1, j-1) //相等，不需要操作
		}
		// 分别对应 新增，删除，替换
		return min(dfs(i, j-1), dfs(i-1, j), dfs(i-1, j-1)) + 1
	}

	return dfs(m-1, n-1)
}
