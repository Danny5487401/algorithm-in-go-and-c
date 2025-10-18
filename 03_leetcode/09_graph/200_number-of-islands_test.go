package _9_graph

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 岛屿数量 https://leetcode.cn/problems/number-of-islands/description/?envType=study-plan-v2&envId=top-100-liked
func TestNumIslands(t *testing.T) {
	convey.Convey("岛屿数量:给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格", t, func() {
		testCase := []struct {
			input  [][]byte
			target int
		}{
			{

				[][]byte{
					{'1', '1', '1', '1', '0'},
					{'1', '1', '0', '1', '0'},
					{'1', '1', '0', '0', '0'},
					{'0', '0', '0', '0', '0'},
				}, 1,
			},

			{

				[][]byte{
					{'1', '1', '0', '0', '0'},
					{'1', '1', '0', '0', '0'},
					{'0', '0', '1', '0', '0'},
					{'0', '0', '0', '1', '1'},
				}, 3,
			},
		}

		for _, tst := range testCase {
			rsp := numIslands(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

// 找到 1 意味着发现了一个新的岛，继续插上旗子🚩，变成 2
// 一旦我们发现 (i,j) 是 1，就从 (i,j) 开始，DFS 这个岛。
// 每一步可以往左右上下四个方向走，也就是 (i,j−1),(i,j+1),(i−1,j),(i+1,j)
func numIslands(grid [][]byte) (ans int) {

	m, n := len(grid), len(grid[0]) // 求行和列

	var dfs func(i, j int)
	dfs = func(i, j int) {
		// 如果 (i,j) 出界，或者 (i,j) 是水，或者 (i,j) 已经插上了旗子 🚩，就不再继续往下递归
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != '1' {
			return

		}
		grid[i][j] = '2' // 已经插上了旗子, 也可以把访问过的 grid[i][j] 改成 0,改成除了 1 以外的任何值都行
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}
	for i, row := range grid {
		for j, col := range row {
			if col == '1' { // 找到岛屿
				dfs(i, j)
				ans++
			}
		}
	}
	return

}
