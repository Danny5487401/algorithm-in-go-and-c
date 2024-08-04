package _5_binary_tree

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 二叉树的最大深度 https://leetcode.cn/problems/maximum-depth-of-binary-tree/?envType=study-plan-v2&envId=top-100-liked

func TestMaxDepth(t *testing.T) {
	convey.Convey("二叉树的最大深度", t, func() {
		testCase := []struct {
			input  *TreeNode
			target int
		}{
			{
				CreateTreeByArray([]int{3, 9, 20, 0, 0, 15, 7}),
				3,
			},
			{
				CreateTreeByArray([]int{1, 0, 2}),
				2,
			},
		}

		for _, tst := range testCase {
			rsp := maxDepth2(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

// 方式一：传递节点
func maxDepth(root *TreeNode) int {
	if root == nil {
		// 边界条件
		return 0
	}
	// 递归 DFS max(leftDepth, rightDepth) + 1
	// 左子树和右子树的最大深度 l 和 r
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	return max(leftDepth, rightDepth) + 1

}

// 方法二：用全局变量保存最大值,传递节点个数
func maxDepth2(root *TreeNode) int {
	var ans = 0
	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		// 每次到一个节点 +1
		depth++
		ans = max(ans, depth)

		// 递归左右子树
		dfs(node.Left, depth)
		dfs(node.Right, depth)
	}
	dfs(root, 0)
	return ans

}
