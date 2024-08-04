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
				Ints2TreeNode([]int{4, 2, 7, 1, 3, 6, 9}),
				Ints2TreeNode([]int{4, 7, 2, 9, 6, 3, 1}),
			},

			{
				Ints2TreeNode([]int{2, 1, 3}),
				Ints2TreeNode([]int{2, 3, 1}),
			},

			{
				Ints2TreeNode([]int{}),
				Ints2TreeNode([]int{}),
			},
		}

		for _, tst := range testCase {
			rsp := invertTree(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func invertTree(root *TreeNode) *TreeNode {
	//对于根节点，它的左右儿子必须交换，即左儿子变成右儿子，右儿子变成左儿子。
	//对于根节点的左右子树，也需要翻转其内部节点。
	if root == nil {
		return nil
	}
	// left 然后 right, 代表从左下面往上
	left := invertTree(root.Left)   // 获取到左子树翻转后的结果 left
	right := invertTree(root.Right) // 获取到右子树翻转后的结果 right

	// 交换左右儿子
	root.Left = right
	root.Right = left
	return root
}
