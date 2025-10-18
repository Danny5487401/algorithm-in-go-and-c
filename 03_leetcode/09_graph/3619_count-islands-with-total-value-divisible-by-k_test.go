package _9_graph

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 总价值可以被 K 整除的岛屿数目 https://leetcode.cn/problems/count-islands-with-total-value-divisible-by-k/description/
func TestCountIslands(t *testing.T) {
	convey.Convey("返回总价值可以被 k 整除 的岛屿数量。", t, func() {
		testCase := []struct {
			input  [][]int
			k      int
			target int
		}{
			{

				[][]int{
					{0, 2, 1, 0, 0},
					{0, 5, 0, 0, 5},
					{0, 0, 1, 0, 0},
					{0, 1, 4, 7, 0},
					{0, 2, 0, 0, 8},
				}, 5, 2,
			},
			{

				[][]int{
					{3, 0, 3, 0},
					{0, 3, 0, 3},
					{3, 0, 3, 0},
				}, 3, 6,
			},
		}

		for _, tst := range testCase {
			rsp := countIslands(tst.input, tst.k)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func countIslands(grid [][]int, k int) int {

	var dir = []struct{ x, y int }{
		{-1, 0},
		{1, 0},
		{0, 1},
		{0, -1},
	}

	m, n := len(grid), len(grid[0])

	var dfs func(x, y int) int

	dfs = func(x, y int) int {
		var area = grid[x][y]
		grid[x][y] = 0
		for _, d := range dir {
			i, j := x+d.x, y+d.y
			if i >= 0 && i < m && j >= 0 && j < n && grid[i][j] > 0 {
				area += dfs(i, j)
			}
		}
		return area
	}
	var ans int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] > 0 && dfs(i, j)%k == 0 {
				ans++
			}
		}
	}
	return ans
}
