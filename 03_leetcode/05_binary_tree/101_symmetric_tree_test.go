package _5_binary_tree

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 对称二叉树 https://leetcode.cn/problems/symmetric-tree/?envType=study-plan-v2&envId=top-100-liked

func TestIsSymmetric(t *testing.T) {
	convey.Convey("对称二叉树", t, func() {
		testCase := []struct {
			root   *TreeNode
			target bool
		}{
			{
				&TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 3,
						},
						Right: &TreeNode{
							Val: 4,
						},
					},
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 4,
						},
						Right: &TreeNode{
							Val: 3,
						},
					},
				},
				true,
			},
			{
				&TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Right: &TreeNode{
							Val: 3,
						},
					},
					Right: &TreeNode{
						Val: 2,
						Right: &TreeNode{
							Val: 3,
						},
					},
				},
				false,
			},
		}

		for _, tst := range testCase {
			rsp := isSymmetric(tst.root)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	// 平分根节点，就与 100. 相同的树 相似：https://leetcode.cn/problems/same-tree/
	return isSymmetricTree(root.Left, root.Right)
}

func isSymmetricTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	// 判断依据 左子树=右子树
	return p.Val == q.Val && isSymmetricTree(p.Left, q.Right) && isSymmetricTree(p.Right, q.Left)
}
