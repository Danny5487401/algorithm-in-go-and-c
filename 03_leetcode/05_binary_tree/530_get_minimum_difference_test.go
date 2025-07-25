package _5_binary_tree

import (
	"github.com/smartystreets/goconvey/convey"
	"math"
	"testing"
)

// 二叉搜索树的最小绝对差 https://leetcode.cn/problems/minimum-absolute-difference-in-bst/description/
func TestGetMinimumDifference(t *testing.T) {
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
		}

		for _, tst := range testCase {
			rsp := getMinimumDifference(tst.root)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func getMinimumDifference(root *TreeNode) int {
	// 绝对值的最小值，答案一定为相邻两个元素之差的最小值，

	pre := -1 // 0 <= Node.val <= 105
	var ans = math.MaxInt
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		// 中序遍历：左子树 → 根 → 右子树

		cur := node.Val
		dfs(node.Left)
		if pre != -1 && cur-pre < ans {
			ans = cur - pre
		}
		pre = node.Val
		dfs(node.Right)
	}

	dfs(root)

	return ans
}
