package _5_binary_tree

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 二叉搜索树中的搜索 https://leetcode.cn/problems/search-in-a-binary-search-tree/
func TestSearchBST(t *testing.T) {
	convey.Convey("二叉搜索树中的搜索:在 BST 中找到节点值等于 val 的节点, 返回以该节点为根的子树", t, func() {
		testCase := []struct {
			root   *TreeNode
			val    int
			target *TreeNode
		}{
			{
				Ints2TreeNode([]int{4, 2, 7, 1, 3}),
				2,
				Ints2TreeNode([]int{2, 1, 3}),
			},
			{
				Ints2TreeNode([]int{4, 2, 7, 1, 3}),
				5,
				Ints2TreeNode([]int{}),
			},
		}

		for _, tst := range testCase {
			rsp := searchBST(tst.root, tst.val)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

// 中序遍历
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}

	if root.Val > val {
		// 在左侧
		return searchBST(root.Left, val)
	} else {
		return searchBST(root.Right, val)
	}

}
