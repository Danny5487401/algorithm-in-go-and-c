package _5_binary_tree

import (
	"github.com/smartystreets/goconvey/convey"
	"math"
	"testing"
)

// 验证二叉搜索树 https://leetcode.cn/problems/validate-binary-search-tree/?envType=study-plan-v2&envId=top-100-liked

// 三种做法
func TestIsValidBST(t *testing.T) {
	convey.Convey("验证二叉搜索树", t, func() {
		testCase := []struct {
			input  *TreeNode
			target bool
		}{
			{
				&TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				true,
			},
			{
				&TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 3,
						},
						Right: &TreeNode{
							Val: 6,
						},
					},
				},
				false,
			},
		}

		for _, tst := range testCase {
			rsp := isValidBST(tst.input)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

// 第一种做法：前序遍历
// 对于每个节点，还需要提供开闭区间的范围。
// 向左：右边界修改
// 向右：左边界修改

func isValidBST(root *TreeNode) bool {
	// 根节点范围 :负无穷 正无穷 (-∞，+∞)

	return checkValidBST(root, math.MinInt, math.MaxInt)
}

func checkValidBST(node *TreeNode, left, right int) bool {
	// 边界条件
	if node == nil {
		return true
	}
	cur := node.Val
	return left < cur && cur < right && checkValidBST(node.Left, left, cur) && checkValidBST(node.Right, cur, right)
}

// 第二种做法：中序遍历
// 基于方法一中提及的性质，我们可以进一步知道二叉搜索树「中序遍历」得到的值构成的序列一定是升序的，这启示我们在中序遍历的时候实时检查当前节点的值是否大于前一个中序遍历到的节点的值即可

func isValidBST2(root *TreeNode) bool {
	// 第一个节点没有上一个节点
	pre := math.MinInt

	var dfs func(*TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		// 检查左子树
		// 当前节点 < 上一个节点
		if !dfs(node.Left) || node.Val <= pre {
			return false
		}
		pre = node.Val
		return dfs(node.Right)
	}
	return dfs(root)

}
