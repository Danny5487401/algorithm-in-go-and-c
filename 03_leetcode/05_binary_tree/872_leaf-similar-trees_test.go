package _5_binary_tree

import (
	"github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)

// 叶子相似的树 https://leetcode.cn/problems/leaf-similar-trees/
func TestLeafSimilar(t *testing.T) {
	convey.Convey("二叉树的中序遍历", t, func() {
		/*
			- 中序遍历: [ [左子树的前序遍历结果], 根节点,[右子树的前序遍历结果] ]
		*/
		testCase := []struct {
			input1 *TreeNode
			input2 *TreeNode
			target bool
		}{
			{
				Ints2TreeNode([]int{3, 5, 1, 6, 2, 9, 8, NULL, NULL, 7, 4}),
				Ints2TreeNode([]int{3, 5, 1, 6, 7, 4, 2, NULL, NULL, NULL, NULL, NULL, NULL, 9, 8}),
				true,
			},
			{
				Ints2TreeNode([]int{1, 2, 3}),
				Ints2TreeNode([]int{1, 3, 2}),
				false,
			},
		}

		for _, tst := range testCase {
			rsp := leafSimilar(tst.input1, tst.input2)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {

	var res []int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root.Left == nil && root.Right == nil {
			res = append(res, root.Val)
			return
		}
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root1)
	res1 := res
	res = []int{}
	dfs(root2)
	return reflect.DeepEqual(res, res1)

}
