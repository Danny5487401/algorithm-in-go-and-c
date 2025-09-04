package _5_binary_tree

import (
	"github.com/smartystreets/goconvey/convey"
	"math"

	"testing"
)

// 统计二叉树中好节点的数目 https://leetcode.cn/problems/count-good-nodes-in-binary-tree/description/
func TestGoodNodes(t *testing.T) {
	convey.Convey("「好节点」X 定义为：从根到该节点 X 所经过的节点中，没有任何节点的值大于 X 的值", t, func() {
		testCase := []struct {
			input  *TreeNode
			target int
		}{
			{
				/*
					根节点 (3) 永远是个好节点。
					节点 4 -> (3,4) 是路径中的最大值。
					节点 5 -> (3,4,5) 是路径中的最大值。
					节点 3 -> (3,1,3) 是路径中的最大值。
				*/
				Ints2TreeNode([]int{3, 1, 4, 3, NULL, 1, 5}),
				4,
			},
			{
				Ints2TreeNode([]int{3, 3, NULL, 4, 2}),
				3,
			},
			{
				Ints2TreeNode([]int{1}),
				1,
			},
		}

		for _, tst := range testCase {
			rsp := goodNodes(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func goodNodes(root *TreeNode) int {
	var dfs func(node *TreeNode, mx int) int // 维护一个参数 mx 表示从根节点到当前节点之前，路径上的最大节点值
	dfs = func(node *TreeNode, mx int) int {
		if node == nil {
			return 0
		}
		// 获取左子树的好节点个数 left
		left := dfs(node.Left, max(mx, node.Val))
		// 获取右子树的好节点个数 right
		right := dfs(node.Right, max(mx, node.Val))
		if mx <= node.Val { // 如果当前节点是好节点
			return left + right + 1
		}
		return left + right
	}
	return dfs(root, math.MinInt32)
}
