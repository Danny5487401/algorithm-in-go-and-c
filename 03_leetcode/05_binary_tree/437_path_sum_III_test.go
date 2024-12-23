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

// 解题思路: 同 560 和为 K 的子数组,使用前缀和+哈希表
func pathSum(root *TreeNode, targetSum int) (ans int) {
	cnt := map[int]int{0: 1}
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
