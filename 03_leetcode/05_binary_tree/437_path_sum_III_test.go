package _5_binary_tree

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 路径总和III https://leetcode.cn/problems/path-sum-iii/description/?envType=study-plan-v2&envId=top-100-liked
func TestPathSum(t *testing.T) {
	convey.Convey("路径总和III:路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下(只能从父节点到子节点）", t, func() {
		testCase := []struct {
			input  *TreeNode
			sum    int
			target int
		}{
			{
				Ints2TreeNode([]int{10, 5, -3, 3, 2, NULL, 11, 3, -2, NULL, 1}),
				8,
				3,
			},
			{
				Ints2TreeNode([]int{5, 4, 8, 11, NULL, 13, 4, 7, 2, NULL, NULL, 5, 1}),
				22,
				3,
			},
		}

		for _, tst := range testCase {
			rsp := pathSum(tst.input, tst.sum)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

// 如果二叉树是一条链，本题就和 560. 和为 K 的子数组 完全一样了：统计有多少个非空连续子数组的元素和恰好等于 targetSum。
// 解题思路: 同 560 和为 K 的子数组,使用前缀和+哈希表
func pathSum(root *TreeNode, targetSum int) (ans int) {
	// 二叉树的节点个数的范围是 [0,1000]

	// 这里的 0 相当于前缀和数组中的 s[0]=0。举个最简单的例子，根节点值为 1，targetSum=1。如果不把 0 加到哈希表中，按照我们的算法，没法算出这里有 1 条符合要求的路径。也可以这样理解，要想把任意路径和都表示成两个前缀和的差，必须添加一个 0，否则当路径是前缀时（从根节点开始的路径），没法减去一个数
	cnt := map[int]int{0: 1} // 计算出现的次数
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}
		sum += node.Val
		ans += cnt[sum-targetSum]
		cnt[sum]++
		dfs(node.Left, sum)
		dfs(node.Right, sum)
		cnt[sum]-- // 恢复现场
	}
	dfs(root, 0)
	return ans
}
