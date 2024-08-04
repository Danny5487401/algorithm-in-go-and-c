package _5_binary_tree

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 路径总和
func TestHasPathSum(t *testing.T) {
	convey.Convey("路径总和", t, func() {
		testCase := []struct {
			input  *TreeNode
			sum    int
			target bool
		}{
			{
				CreateTreeByArray([]int{5, 4, 8, 11, 0, 13, 4, 7, 2, 0, 0, 0, 1}),
				22,
				true,
			},
			{
				CreateTreeByArray([]int{1, 2, 3}),
				5,
				false,
			},
			{
				CreateTreeByArray([]int{}),
				0,
				false,
			},
		}

		for _, tst := range testCase {
			rsp := hasPathSum(tst.input, tst.sum)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func hasPathSum(root *TreeNode, targetSum int) bool {

	if root == nil {
		return false // Since the tree is empty, there are no root-to-leaf paths
	}

	targetSum -= root.Val

	if root.Left == nil && root.Right == nil { // 叶子节点
		return targetSum == 0
	}
	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}
