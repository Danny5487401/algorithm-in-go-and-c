package _5_binary_tree

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 二叉搜索树的范围和  https://leetcode.cn/problems/range-sum-of-bst/
func TestRangeSumBST(t *testing.T) {
	convey.Convey("二叉搜索树的范围和:返回值位于范围 [low, high] 之间的所有结点的值的和", t, func() {
		testCase := []struct {
			root   *TreeNode
			low    int
			high   int
			target int
		}{
			{
				Ints2TreeNode([]int{10, 5, 15, 3, 7, NULL, 18}),
				7,
				15,
				32,
			},
			{
				Ints2TreeNode([]int{10, 5, 15, 3, 7, 13, 18, 1, NULL, 6}),
				6,
				10,
				23,
			},
		}

		for _, tst := range testCase {
			rsp := rangeSumBST(tst.root, tst.low, tst.high)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil { // root 节点为空
		return 0
	}

	// 根节点的值为 x
	x := root.Val

	if x < low { // 左子树没有节点在范围内，只需递归右子树
		// 左子树的所有节点值都小于 x，从而小于 low，所以也不在 [low,high] 范围内，我们只需计算右子树的在 [low,high] 范围内的节点值之和。
		return rangeSumBST(root.Right, low, high)

	} else if x > high { //  右子树没有节点在范围内，只需递归左子树
		// 右子树的所有节点值都大于 x，从而大于 high，所以也不在 [low,high] 范围内，我们只需计算左子树的在 [low,high] 范围内的节点值之和。

		return rangeSumBST(root.Left, low, high)

	} else {
		// 累加 x、左子树的在 [low,high] 范围内的节点值之和、右子树的在 [low,high] 范围内的节点值之和，这三部分的和作为答案
		return x + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
	}

}
