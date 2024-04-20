package _5_binary_tree

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

// 翻转二叉树: https://leetcode.cn/problems/invert-binary-tree/?envType=study-plan-v2&envId=top-100-liked

func TestInvertTree(t *testing.T) {
	convey.Convey("翻转二叉树", t, func() {
		testCase := []struct {
			input  *TreeNode
			target *TreeNode
		}{
			{
				&TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 1,
						},
						Right: &TreeNode{
							Val: 3,
						},
					},
					Right: &TreeNode{
						Val: 7,
						Left: &TreeNode{
							Val: 6,
						},
						Right: &TreeNode{
							Val: 9,
						},
					},
				},
				&TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 7,
						Left: &TreeNode{
							Val: 9,
						},
						Right: &TreeNode{
							Val: 6,
						},
					},
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 3,
						},
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
			},
		}

		for _, tst := range testCase {
			rsp := invertTree(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func invertTree(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}
	// left 然后 right, 代表从左下面往上
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}
