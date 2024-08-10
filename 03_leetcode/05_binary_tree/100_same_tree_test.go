package _5_binary_tree

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 相同的树 https://leetcode.cn/problems/same-tree/
func TestIsSameTree(t *testing.T) {
	convey.Convey("相同的树", t, func() {
		testCase := []struct {
			left   *TreeNode
			right  *TreeNode
			target bool
		}{
			{
				Ints2TreeNode([]int{1, 2, 3}),
				Ints2TreeNode([]int{1, 2, 3}),
				true,
			},
			{
				Ints2TreeNode([]int{1, 2}),
				Ints2TreeNode([]int{1, 0, 2}),
				false,
			},
		}

		for _, tst := range testCase {
			rsp := isSameTree(tst.left, tst.right)
			convey.So(rsp, convey.ShouldEqual, tst.target)
		}
	})

}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	// 边界条件
	if p == nil || q == nil {
		return p == q
	}
	// 判断依据是左子树=左子树
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
