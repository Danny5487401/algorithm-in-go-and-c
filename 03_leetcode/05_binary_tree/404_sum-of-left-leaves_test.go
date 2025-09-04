package _5_binary_tree

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 左叶子之和 https://leetcode.cn/problems/sum-of-left-leaves/description/
func TestSumOfLeftLeaves(t *testing.T) {
	convey.Convey("左叶子之和", t, func() {
		testCase := []struct {
			input  *TreeNode
			target int
		}{
			{
				// 这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24
				Ints2TreeNode([]int{3, 9, 20, NULL, NULL, 15, 7}),
				24,
			},
			{
				Ints2TreeNode([]int{1}),
				0,
			},
			{
				Ints2TreeNode([]int{1, NULL, 2}),
				0,
			},
		}

		for _, tst := range testCase {
			rsp := sumOfLeftLeaves(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func sumOfLeftLeaves(root *TreeNode) int {

	if root == nil {
		return 0
	}
	sum := sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil { // 当前节点的左儿子是叶子
		sum += root.Left.Val
	}
	return sum

}
