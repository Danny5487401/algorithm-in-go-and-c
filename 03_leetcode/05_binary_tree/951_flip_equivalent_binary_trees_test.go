package _5_binary_tree

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 翻转等价二叉树 https://leetcode.cn/problems/flip-equivalent-binary-trees/
func TestFlipEquiv(t *testing.T) {
	convey.Convey("翻转等价二叉树:选择任意节点，然后交换它的左子树和右子树 ", t, func() {
		testCase := []struct {
			root1  *TreeNode
			root2  *TreeNode
			target bool
		}{
			{
				Ints2TreeNode([]int{1, 2, 3, 4, 5, 6, NULL, NULL, NULL, 7, 8}),
				Ints2TreeNode([]int{1, 3, 2, NULL, 6, 4, 5, NULL, NULL, NULL, NULL, 8, 7}),
				true,
			},
			{
				Ints2TreeNode([]int{}),
				Ints2TreeNode([]int{}),
				true,
			},
			{
				Ints2TreeNode([]int{}),
				Ints2TreeNode([]int{1}),
				false,
			},
		}

		for _, tst := range testCase {
			rsp := flipEquiv(tst.root1, tst.root2)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if (root1 == nil && root2 != nil) || (root2 == nil && root1 != nil) {
		return false
	}

	if root1.Val == root2.Val {
		return (flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right)) || // 不翻转
			(flipEquiv(root1.Left, root2.Right) && flipEquiv(root1.Right, root2.Left)) // 翻转
	}
	return false

}
