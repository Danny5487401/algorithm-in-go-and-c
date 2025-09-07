package _9_graph

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 课程表 II https://leetcode.cn/problems/course-schedule/description/
func TestFindOrder(t *testing.T) {
	convey.Convey("课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 . 图有环代表完成不了."+
		"如果不可能完成所有课程，返回 一个空数组 。", t, func() {
		testCase := []struct {
			input         int
			prerequisites [][]int
			target        []int
		}{
			// 总共有 2 门课程。要学习课程 1，你需要先完成课程 0。
			{

				2,
				[][]int{
					{1, 0},
				},
				[]int{0, 1},
			},
			// 总共有 4 门课程。要学习课程 3，你应该先完成课程 1 和课程 2。并且课程 1 和课程 2 都应该排在课程 0 之后。
			{
				4,
				[][]int{
					{1, 0},
					{2, 0},
					{3, 1},
					{3, 2},
				},
				[]int{0, 2, 1, 3},
			},

			{
				1,
				[][]int{},
				[]int{0},
			},

			{
				3,
				[][]int{
					{0, 2},
					{1, 2},
					{2, 0},
				},
				[]int{},
			},
		}

		for _, tst := range testCase {
			rsp := findOrder(tst.input, tst.prerequisites)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

/*
对于图中的任意一个节点，它在搜索的过程中有三种状态，即：

「未搜索」：我们还没有搜索到这个节点；

「搜索中」：我们搜索过这个节点，但还没有回溯到该节点，即该节点还没有入栈，还有相邻的节点没有搜索完成）；

「已完成」：我们搜索过并且回溯过这个节点，即该节点已经入栈，并且所有该节点的相邻节点都出现在栈的更底部的位置，满足拓扑排序的要求



我们将当前搜索的节点 u 标记为「搜索中」，遍历该节点的每一个相邻节点 v：

如果 v 为「未搜索」，那么我们开始搜索 v，待搜索完成回溯到 u；

如果 v 为「搜索中」，那么我们就找到了图中的一个环，因此是不存在拓扑排序的；

如果 v 为「已完成」，那么说明 v 已经在栈中了，而 u 还不在栈中，因此 u 无论何时入栈都不会影响到 (u,v) 之前的拓扑关系，以及不用进行任何操作。

当 u 的所有相邻节点都为「已完成」时，我们将 u 放入栈中，并将其标记为「已完成」。

在整个深度优先搜索的过程结束后，如果我们没有找到图中的环，那么栈中存储这所有的 n 个节点，从栈顶到栈底的顺序即为一种拓扑排序。

*/

// 深度优先搜索
func findOrder(numCourses int, prerequisites [][]int) []int {

	// 每个点的边
	var edge = make([][]int, numCourses)
	visited := make([]int, numCourses)

	// 栈记录依赖顺序
	result := []int{}

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
	// 记录边
	for _, info := range prerequisites {
		edge[info[1]] = append(edge[info[1]], info[0])
	}

	// 开始修课程

	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	if !valid {
		return []int{}
	}

	// 需要的是依赖在数组前面
	for i := 0; i < len(result)/2; i++ {
		result[i], result[numCourses-i-1] = result[numCourses-i-1], result[i]
	}
	return result

}
