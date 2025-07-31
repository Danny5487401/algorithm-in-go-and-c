package _5_binary_tree

import (
	"github.com/smartystreets/goconvey/convey"
	"math"
	"testing"
)

// 二叉搜索树节点最小距离 https://leetcode.cn/problems/minimum-distance-between-bst-nodes/description/
func TestMinDiffInBST(t *testing.T) {
	convey.Convey("二叉搜索树的最小绝对差:树中任意两不同节点值之间的最小差值,中序遍历 ", t, func() {
		testCase := []struct {
			root   *TreeNode
			target int
		}{
			{
				// 把数字排序后，得到 [1,2,3,4,6]。
				Ints2TreeNode([]int{4, 2, 6, 1, 3}),
				1,
			},
			{
				Ints2TreeNode([]int{1, 0, 48, NULL, NULL, 12, 49}),
				1,
			},
			{
				Ints2TreeNode([]int{99, 84, NULL, 27, NULL, 1, 53}),
				1,
			},
		}

		for _, tst := range testCase {
			rsp := minDiffInBST(tst.root)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func minDiffInBST(root *TreeNode) int {
	ans := math.MaxInt64
	pre := math.MinInt64 / 2
	var dfs func(root *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		ans = min(ans, node.Val-pre)
		pre = node.Val
		dfs(node.Right)

	}
	dfs(root)
	return ans
}
