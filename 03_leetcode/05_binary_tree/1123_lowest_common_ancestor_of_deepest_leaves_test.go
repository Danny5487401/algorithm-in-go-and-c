package _5_binary_tree

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 最深叶节点的最近公共祖先 https://leetcode.cn/problems/lowest-common-ancestor-of-deepest-leaves/
func TestLcaDeepestLeaves(t *testing.T) {
	convey.Convey("最深叶节点的最近公共祖先 :", t, func() {
		testCase := []struct {
			input  *TreeNode
			target *TreeNode
		}{
			{
				Ints2TreeNode([]int{3, 5, 1, 6, 2, 0, 8, NULL, NULL, 7, 4}),
				Ints2TreeNode([]int{2, 7, 4}),
			},

			{
				Ints2TreeNode([]int{1}),
				Ints2TreeNode([]int{1}),
			},

			{
				Ints2TreeNode([]int{0, 1, 3, NULL, 2}),
				Ints2TreeNode([]int{2}),
			},
		}

		for _, tst := range testCase {
			rsp := lcaDeepestLeaves(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	// 如果左子树的最大深度比右子树的大，那么最深叶结点就只在左子树中，所以最近公共祖先也只在左子树中。

	/*
		1. 递归这棵二叉树，同时维护全局最大深度 maxDepth。
		2. 在「递」的时候往下传 depth，用来表示当前节点的深度。
		3. 在「归」的时候往上传当前子树最深叶节点的深度。
		4. 设左子树最深叶节点的深度为 leftMaxDepth，右子树最深叶节点的深度为 rightMaxDepth。如果 leftMaxDepth=rightMaxDepth=maxDepth，那么更新答案为当前节点。注意这并不代表我们找到了答案，如果后面发现了更深的叶节点，那么答案还会更新。

	*/
	var maxDepth = -1
	var ans *TreeNode
	var dfs func(node *TreeNode, depth int) int

	dfs = func(node *TreeNode, depth int) int {
		if node == nil {
			maxDepth = max(maxDepth, depth)
			return depth
		}
		leftMaxDepth := dfs(node.Left, depth+1)
		rightMaxDepth := dfs(node.Right, depth+1)
		if leftMaxDepth == maxDepth && rightMaxDepth == maxDepth {
			ans = node
		}
		return max(leftMaxDepth, rightMaxDepth)

	}

	dfs(root, 0)
	return ans
}
