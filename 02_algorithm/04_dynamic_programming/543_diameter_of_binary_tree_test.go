package _4_dynamic_programming

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

// 二叉树的直径 https://leetcode.cn/problems/diameter-of-binary-tree/description/

func TestDiameterOfBinaryTree(t *testing.T) {
	convey.Convey("二叉树的直径：两点最大距离", t, func() {
		testCase := []struct {
			input  *TreeNode
			target int
		}{
			{
				Ints2TreeNode([]int{1, 2, 3, 4, 5}),
				3,
			},
			{
				Ints2TreeNode([]int{1, 2}),
				1,
			},
		}

		for _, tst := range testCase {
			rsp := diameterOfBinaryTree1(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func diameterOfBinaryTree(root *TreeNode) int {

	var dfs func(node *TreeNode) int
	var ans int
	dfs = func(node *TreeNode) int {
		if node == nil {
			// 当叶子节点时
			return -1
		}
		leftLength := dfs(node.Left)
		rightLength := dfs(node.Right)
		// 当前拐弯的直接长度 = 左子路的最长链 + 右子路的最长链 + 2
		ans = max(ans, leftLength+rightLength+2)

		// 返回父节点节点 max(leftLength, rightLength) + 1
		return max(leftLength, rightLength) + 1
	}
	dfs(root)
	return ans
}

func diameterOfBinaryTree1(root *TreeNode) int {

	var dfs func(node *TreeNode) int
	var ans int
	dfs = func(node *TreeNode) int {
		if node == nil {
			// 当叶子节点时
			return -1
		}
		leftLength := dfs(node.Left) + 1
		rightLength := dfs(node.Right) + 1
		// 当前拐弯的直接长度 = 左子路的最长链 + 右子路的最长链
		ans = max(ans, leftLength+rightLength)

		// 返回父节点节点 max(leftLength, rightLength)
		return max(leftLength, rightLength)
	}
	dfs(root)
	return ans
}
