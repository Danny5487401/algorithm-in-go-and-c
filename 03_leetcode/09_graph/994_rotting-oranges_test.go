package _9_graph

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 腐烂的橘子 https://leetcode.cn/problems/rotting-oranges/description/?envType=study-plan-v2&envId=top-100-liked
func TestOrangesRotting(t *testing.T) {
	convey.Convey("腐烂的橘子", t, func() {
		testCase := []struct {
			input  [][]int
			target int
		}{
			{

				[][]int{
					{2, 1, 1},
					{1, 1, 0},
					{0, 1, 1},
				}, 4,
			},

			{

				[][]int{
					{2, 1, 1},
					{0, 1, 1},
					{1, 0, 1},
				}, -1,
			},

			{

				[][]int{
					{0, 2},
				}, 0,
			},
		}

		for _, tst := range testCase {
			rsp := orangesRotting(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

type pair struct{ x, y int }

var directions = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 四方向

func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	// 为了判断是否有永远不会腐烂的橘子，我们可以统计初始新鲜橘子的个数 fresh
	fresh := 0

	// 统计所有初始就腐烂的橘子的位置
	q := []pair{}
	for i, row := range grid {
		for j, x := range row {
			if x == 1 {
				fresh++ // 统计新鲜橘子个数
			} else if x == 2 {
				q = append(q, pair{i, j}) // 一开始就腐烂的橘子
			}
		}
	}

	ans := 0
	for fresh > 0 && len(q) > 0 {
		ans++ // 经过一分钟
		tmp := q
		q = []pair{}
		for _, p := range tmp { // 已经腐烂的橘子
			for _, d := range directions { // 四方向
				i, j := p.x+d.x, p.y+d.y
				if 0 <= i && i < m && 0 <= j && j < n && grid[i][j] == 1 { // 新鲜橘子
					fresh--
					grid[i][j] = 2 // 变成腐烂橘子
					q = append(q, pair{i, j})
				}
			}
		}
	}

	if fresh > 0 { // 还有没有腐烂的橘子
		return -1
	}
	return ans
}
