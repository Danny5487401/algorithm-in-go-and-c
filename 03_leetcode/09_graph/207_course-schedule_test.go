package _9_graph

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

// 课程表 https://leetcode.cn/problems/course-schedule/description/
func TestCanFinish(t *testing.T) {
	convey.Convey("课程对 [0, 1] 表示：1-->0, 想要学习课程 0 ，你需要先完成课程 1. 图有环代表完成不了", t, func() {
		testCase := []struct {
			input         int
			prerequisites [][]int
			target        bool
		}{
			{

				2,
				[][]int{
					{1, 0},
				},
				true,
			},

			{
				2,
				[][]int{
					{1, 0},
					{0, 1},
				},
				false,
			},
		}

		for _, tst := range testCase {
			rsp := canFinish(tst.input, tst.prerequisites)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

/*
我们将每一门课看成一个节点；

如果想要学习课程 A 之前必须完成课程 B，那么我们从 B 到 A 连接一条有向边。这样以来，在拓扑排序中，B 一定出现在 A 的前面。

对于每个节点 x，都定义三种颜色值（状态值）：

0：节点 x 尚未被访问到。
1：节点 x 正在访问中，dfs(x) 尚未结束。
2：节点 x 已经完全访问完毕，dfs(x) 已返回。
*/
func canFinish(numCourses int, prerequisites [][]int) bool {
	// 每个点的边
	var edge = make([][]int, numCourses)

	// 创建长为 numCourses 的访问数组
	visited := make([]int, numCourses)

	// 栈记录依赖顺序
	result := []int{}

	// 1 建图：把每个 prerequisites[i]=[a,b] 看成一条有向边 b→a，构建一个有向图 g。
	// 记录边
	for _, info := range prerequisites {
		edge[info[1]] = append(edge[info[1]], info[0])
	}

	var valid bool = true // 有效,无环

	var dfs func(u int)
	dfs = func(u int) {
		visited[u] = 1 // 搜索中
		for _, v := range edge[u] {
			if visited[v] == 0 { // 「未搜索」
				dfs(v)
				if !valid {
					return
				}

			} else if visited[v] == 1 {
				valid = false // 存在环
				return
			}

		}
		visited[u] = 2             //  已完成
		result = append(result, u) // 默认最先入站的是无依赖, 最后是依赖

	}

	// 开始修课程
	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	return valid
}
