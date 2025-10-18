package _9_graph

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 岛屿的最大面积 https://leetcode.cn/problems/max-area-of-island/description/
func TestMaxAreaOfIsland(t *testing.T) {
	convey.Convey("岛屿数量:给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格", t, func() {
		testCase := []struct {
			input  [][]int
			target int
		}{
			{

				[][]int{
					{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
					{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
					{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
				}, 6,
			},
			{

				[][]int{
					{1, 1, 0, 0, 0},
					{1, 1, 0, 0, 0},
					{0, 0, 0, 1, 1},
					{0, 0, 0, 1, 1},
				}, 4,
			},
		}

		for _, tst := range testCase {
			rsp := maxAreaOfIsland(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

var dirs = []struct{ x, y int }{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} // 左右上下

func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	// 只有为1才进入
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		var area = 1
		grid[i][j] = 0 // 标记 (i,j) 访问过
		for _, d := range dirs {
			x, y := i+d.x, j+d.y
			if 0 <= x && x < m && 0 <= y && y < n && grid[x][y] > 0 {
				// 把统计岛屿面积的任务交给其他人去处理，自己只需累加其他人统计出来的岛屿面积
				area += dfs(x, y)
			}
		}

		return area
	}
	var ans int
	for i, r := range grid {
		for j, val := range r {
			if val == 1 {
				ans = max(ans, dfs(i, j))
			}
		}
	}
	return ans
}
