package _4_dynamic_programming

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 最长同值路径 https://leetcode.cn/problems/longest-univalue-path/
func TestLongestUnivaluePath(t *testing.T) {
	convey.Convey("最长同值路径: 边数", t, func() {
		testCase := []struct {
			input  *TreeNode
			target int
		}{
			{
				Ints2TreeNode([]int{5, 4, 5, 1, 1, 5}),
				2,
			},
			{
				Ints2TreeNode([]int{1, 4, 5, 4, 4, 5}),
				2,
			},
			{
				Ints2TreeNode([]int{1, NULL, 1, 1, 1, 1, 1, 1}),
				4,
			},
		}

		for _, tst := range testCase {
			rsp := longestUnivaluePath(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func longestUnivaluePath(root *TreeNode) int {
	/*
	   	在 543. 二叉树的直径 的基础上修改：

	      如果当前节点和左子树的值不同，左子树的链长可以视作 0，否则就是左子树的链长。
	      如果当前节点和右子树的值不同，右子树的链长可以视作 0，否则就是右子树的链长；

	*/
	var ans = -1
	if root == nil {
		return 0
	}
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			// 当叶子节点时
			return -1
		}
		leftSum := dfs(node.Left) + 1
		rightSum := dfs(node.Right) + 1
		if node.Left != nil && node.Left.Val != node.Val {
			leftSum = 0
		}
		if node.Right != nil && node.Right.Val != node.Val {
			rightSum = 0
		}
		ans = max(ans, leftSum+rightSum) // 两条链拼成路径
		return max(leftSum, rightSum)    // 当前子树最大链长
	}

	dfs(root)
	return ans
}
