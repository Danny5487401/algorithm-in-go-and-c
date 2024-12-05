package _9_graph

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 岛屿数量 https://leetcode.cn/problems/number-of-islands/description/?envType=study-plan-v2&envId=top-100-liked
func TestMerge(t *testing.T) {
	convey.Convey("岛屿数量", t, func() {
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

func numIslands(grid [][]byte) (ans int) {

	m, n := len(grid), len(grid[0])

	var dfs func(i, j int)
	dfs = func(i, j int) {
		// 如果 (i,j) 出界，或者 (i,j) 是水，或者 (i,j) 已经插上了旗子🚩，就不再继续往下递归
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != '1' {
			return

		}
		grid[i][j] = '2' // 已经插上了旗子
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
