package _5_binary_tree

import (
	"github.com/smartystreets/goconvey/convey"
	"slices"
	"testing"
)

// 二叉树的层序遍历 https://leetcode.cn/problems/binary-tree-level-order-traversal-ii/description/
func TestLevelOrderBottom(t *testing.T) {
	convey.Convey("二叉树的层序遍历II: 自底向上的层序遍历 ", t, func() {
		testCase := []struct {
			input  *TreeNode
			target [][]int
		}{
			{
				Ints2TreeNode([]int{3, 9, 20, NULL, NULL, 15, 7}),
				[][]int{
					{15, 7},
					{9, 20},
					{3},
				},
			},
			{
				Ints2TreeNode([]int{1}),
				[][]int{
					{1},
				},
			},
			{
				Ints2TreeNode([]int{}),
				[][]int{},
			},
		}

		for _, tst := range testCase {
			rsp := levelOrderBottom(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})
}

// 一个队列
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		// node 范围 [0, 2000]
		return [][]int{}
	}
	var ans = make([][]int, 0)
	var queue = []*TreeNode{root}
	for len(queue) > 0 {
		val := make([]int, len(queue))

		for i := range val {
			node := queue[0]
			val[i] = node.Val
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, val)

	}
	slices.Reverse(ans)
	return ans
}
