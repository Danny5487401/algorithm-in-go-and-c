package _4_dynamic_programming

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 相邻字符不同的最长路径 https://leetcode.cn/problems/longest-path-with-different-adjacent-characters/description/
func TestLongestPath(t *testing.T) {
	convey.Convey("相邻字符不同的最长路径", t, func() {
		testCase := []struct {
			parent []int
			s      string
			target int
		}{
			{
				[]int{-1, 0, 0, 1, 1, 2}, "abacbe", 3,
			},
			{

				[]int{-1, 0, 0, 0}, "aabc", 3,
			},
		}

		for _, tst := range testCase {
			rsp := longestPath(tst.parent, tst.s)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func longestPath(parent []int, s string) int {
	// 边是以 parent 的形式
	tree := make([][]int, len(parent))

	// 求出邻居

	length := len(parent)
	for i := 1; i < length; i++ {
		tree[parent[i]] = append(tree[parent[i]], i) // 第 i 个节点的邻居的下标
	}
	var ans int
	var dfs func(int) int
	dfs = func(x int) (maxLen int) {
		for _, y := range tree[x] {
			yLen := dfs(y) + 1
			if s[x] != s[y] {
				ans = max(ans, maxLen+yLen)
				maxLen = max(maxLen, yLen)
			}

		}
		return
	}
	dfs(0)

	// 点的数量 = 边的数量+1
	return ans + 1
}
