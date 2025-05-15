package _5_binary_tree

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// N 叉树的最大深度 https://leetcode.cn/problems/maximum-depth-of-n-ary-tree/description/
func TestMaxDepthOfNAryTree(t *testing.T) {
	convey.Convey("N 叉树的最大深度", t, func() {
		testCase := []struct {
			input  *MulNode
			target int
		}{
			{
				input: &MulNode{
					Val: 1,
					Children: []*MulNode{
						{Val: 3,
							Children: []*MulNode{
								{Val: 5},
								{Val: 6},
							}},
						{Val: 2},
						{Val: 4},
					},
				},
				target: 3,
			},
		}

		for _, tst := range testCase {
			rsp := maxDepthOfNAryTree(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func maxDepthOfNAryTree(root *MulNode) int {

	var dfs func(*MulNode, int) int
	dfs = func(node *MulNode, depth int) int {
		if node == nil {
			return depth
		}

		// 记 N 个子树的最大深度中的最大值为 maxChildDepth，则该 N 叉树的最大深度为 maxChildDepth+1
		maxChildDepth := 0
		for _, child := range node.Children {
			temDepth := dfs(child, depth)
			maxChildDepth = max(maxChildDepth, temDepth)
		}
		return maxChildDepth + 1
	}

	return dfs(root, 0)
}

type MulNode struct {
	Val      int
	Children []*MulNode
}
