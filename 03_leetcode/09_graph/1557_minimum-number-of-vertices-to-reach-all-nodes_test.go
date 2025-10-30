package _9_graph

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 可以到达所有点的最少点数目 https://leetcode.cn/problems/minimum-number-of-vertices-to-reach-all-nodes/description/
func TestFindSmallestSetOfVertices(t *testing.T) {
	convey.Convey("给你一个 有向无环图 ， n 个节点编号为 0 到 n-1 ，以及一个边数组 edges ，其中 edges[i] = [from_i, to_i] 表示一条从点  from_i 到点 to_i 的有向边。"+
		"找到最小的点集使得从这些点出发能到达图中所有点。题目保证解存在且唯一", t, func() {
		testCase := []struct {
			input         int
			prerequisites [][]int
			target        []int
		}{
			{

				6,
				[][]int{
					{0, 1},
					{0, 2},
					{2, 5},
					{3, 4},
					{4, 2},
				},
				// 从单个节点出发无法到达所有节点。从 0 出发我们可以到达 [0,1,2,5] 。从 3 出发我们可以到达 [3,4,2,5] 。所以我们输出 [0,3] 。
				[]int{0, 3},
			},

			{
				5,
				[][]int{
					{0, 1},
					{2, 1},
					{3, 1},
					{1, 4},
					{2, 4},
				},
				[]int{0, 2, 3},
			},

			{
				3,
				[][]int{
					{1, 2},
					{1, 0},
					{0, 2},
				},
				[]int{1},
			},
		}

		for _, tst := range testCase {
			rsp := findSmallestSetOfVertices(tst.input, tst.prerequisites)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	g := make([][]int, n)
	inDeg := make([]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		inDeg[y]++ // 统计 y 的先修课数量
	}

	q := make([]int, 0, n)
	for i, d := range inDeg {
		if d == 0 { // 没有先修课，可以直接上
			q = append(q, i) // 加入学习队列
		}
	}
	return q
}
